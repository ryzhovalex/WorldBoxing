package orwynn

import (
	"encoding/binary"
	"worldboxing/lib/utils"
)

const CodeCannotReferenceInnerTransport = 2
const CodeIncorrectMessageStructure = 3
const CodeTransportClosed = 4
const CodeConnectionClosed = 5
const CodePublishWithoutTargetConnections = 6

type Code = utils.Code
type MessageBody = []byte
type Message struct {
	Code Code
	Body MessageBody
}

func (message *Message) Serialize() []byte {
	codeBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(codeBytes, uint16(message.Code))
	result := make([]byte, 2+len(message.Body))
	result = append(result, codeBytes...)
	result = append(result, message.Body...)
	return result
}

type MessageContext struct {
	Message *Message
	// AuthorConnection of message's author. Set to nil if author is the host.
	AuthorConnection Connection
	// Connections to which the message should be sent to. Only considered for
	// output messages. If not set, message processing is no-op.
	TargetConnections []Connection
}

// Send a response message to the author connection of this message.
// Helper method to grab context's author and use it as a target in Publish.
func (ctx *MessageContext) Answer(code Code, body MessageBody) *utils.Error {
	return Publish(
		code,
		body,
		[]Connection{ctx.AuthorConnection},
	)
}

type Subscriber func(*MessageContext)
type Subscription struct {
	Id         utils.Id
	Code       Code
	Subscriber Subscriber
}

const MessageQueueMaxSize = 65536

type State struct {
	Transports          map[string]Transport
	CodeToSubscriptions map[Code][]*Subscription
	// TODO: Eliminate risk of available subscription id overflow while
	//		 earlier slots became vacant (by unsubscribing).
	NextSubscriptionId utils.Id
}

type UnsubscribeFunction func()

var state *State

func Init(transports map[string]Transport) *utils.Error {
	_, ok := transports["INNER"]
	if ok {
		return utils.NewError(CodeCannotReferenceInnerTransport, "")
	}
	state = &State{
		Transports:          transports,
		CodeToSubscriptions: map[Code][]*Subscription{},
		NextSubscriptionId:  0,
	}
	for _, transport := range transports {
		go acceptTransportConnections(transport)
	}
	return nil
}

func acceptTransportConnections(transport Transport) {
	defer transport.Close()
	for {
		connection, e := transport.Accept()
		if e != nil {
			utils.Log(e.Error())
			break
		}
		go processConnectionInput(connection)
	}
}

func processConnectionInput(connection Connection) {
	defer connection.Close()
	for {
		// if any message processing error happens, the connection is
		// immediately closed
		raw, e := connection.Recv()
		if e != nil {
			utils.Log(e.Error())
			break
		}
		message, e := newMessage(raw)
		if e != nil {
			utils.Log(e.Error())
			break
		}
		messageContext := MessageContext{
			Message:          message,
			AuthorConnection: connection,
		}
		go sendMessageContextToInner(&messageContext)
	}
}

func sendMessageContextToInner(ctx *MessageContext) {
	subs, ok := state.CodeToSubscriptions[ctx.Message.Code]
	if !ok {
		// no subs, skip
		return
	}
	for _, sub := range subs {
		go sub.Subscriber(ctx)
	}
}

func distributeMessageContextToTargetConnections(ctx *MessageContext) {
	targetConnectionsLen := len(ctx.TargetConnections)
	if targetConnectionsLen == 0 {
		return
	}
	if targetConnectionsLen == 1 {
		// in most cases we send to one recipient, so we save a little bit
		// by not going into a loop, but index first connection directly
		go ctx.TargetConnections[0].Send(ctx.Message.Serialize())
		return
	}
	for _, targetConnection := range ctx.TargetConnections {
		go targetConnection.Send(ctx.Message.Serialize())
	}
}

// Spec:
// First two bytes -> int16 code
// Rest is to the body.
func newMessage(raw []byte) (*Message, *utils.Error) {
	if len(raw) < 2 {
		return nil, utils.NewError(CodeIncorrectMessageStructure, "")
	}
	rawCode := raw[:2]
	var code Code
	binary.LittleEndian.PutUint16(rawCode, uint16(code))
	code = int16(binary.LittleEndian.Uint16(rawCode))
	var body = []byte{}
	if len(raw) > 2 {
		body = raw[2:]
	}
	message := Message{
		Code: code,
		Body: body,
	}
	return &message, nil
}

// Publish a message body to the bus.
func Publish(code Code, body MessageBody, targets []Connection) *utils.Error {
	if len(targets) == 0 {
		return utils.NewError(CodePublishWithoutTargetConnections, "")
	}

	message := Message{
		Code: code,
		Body: body,
	}
	ctx := MessageContext{
		Message:           &message,
		AuthorConnection:  nil,
		TargetConnections: targets,
	}
	// no need to create thread here: cheaper would be to iterate over
	// target connections, since in usual case we send to 1 recipient
	distributeMessageContextToTargetConnections(&ctx)
	return nil
}

func Subscribe(code Code, subscriber Subscriber) (func(), *utils.Error) {
	subscriptionId := state.NextSubscriptionId
	state.NextSubscriptionId += 1
	subscription := Subscription{
		Id:         subscriptionId,
		Code:       code,
		Subscriber: subscriber,
	}

	_, ok := state.CodeToSubscriptions[code]
	if !ok {
		state.CodeToSubscriptions[code] = []*Subscription{}
	}
	state.CodeToSubscriptions[code] = append(
		state.CodeToSubscriptions[code],
		&subscription,
	)

	return func() {
		subscriptions, ok := state.CodeToSubscriptions[code]
		if !ok {
			return
		}
		indexToDelete := -1
		for i, subscription := range subscriptions {
			if subscription.Id == subscriptionId {
				indexToDelete = i
				break
			}
		}
		state.CodeToSubscriptions[code] = utils.RemoveFromUnorderedSlice(
			state.CodeToSubscriptions[code],
			indexToDelete,
		)
	}, nil
}
