// Making reactivity.
//
// Keypoints:
// * we have a thread per connection
// * we have a single thread for output queue (may be changed in a future)
// * each subscriber is called in a separate thread
package orwynn

import (
	"encoding/binary"
	"worldboxing/lib/utils"
)

const CodeCannotReferenceInnerTransport = 2
const CodeIncorrectMessageStructure = 3
const CodeTransportClosed = 4
const CodeConnectionClosed = 5

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
	// output messages.
	TargetConnections []Connection
}

// Send a response message to the author connection of this message.
func (ctx *MessageContext) Answer(code Code, body MessageBody) *utils.Error {
	return Publish(
		code,
		body,
		&PublishOptions{
			TargetConnections: []Connection{ctx.AuthorConnection},
		},
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
	// We have a single incoming and outcoming InputQueue for all transports.
	InputQueue          *MessageQueue
	OutputQueue         *MessageQueue
	Transports          map[string]Transport
	CodeToSubscriptions map[Code][]*Subscription
	// TODO: Eliminate risk of available subscription id overflow while
	//		 earlier slots became vacant (by unsubscribing).
	NextSubscriptionId utils.Id
}

type PublishOptions struct {
	// Connections to publish to. Inner transport is used in any case.
	TargetConnections []Connection
}

type UnsubscribeFunction func()

var state *State

func Init(transports map[string]Transport) *utils.Error {
	_, ok := transports["INNER"]
	if ok {
		return utils.NewError(CodeCannotReferenceInnerTransport)
	}
	state = &State{
		InputQueue:          NewMessageQueue(MessageQueueMaxSize),
		OutputQueue:         NewMessageQueue(MessageQueueMaxSize),
		Transports:          transports,
		CodeToSubscriptions: map[Code][]*Subscription{},
		NextSubscriptionId:  0,
	}
	go processOutputQueue()
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
		state.InputQueue.Enqueue(&messageContext)
	}
}

// How long to sleep if output queue is empty. So it will take this time for
// the server for awake after an empty queue.
const OutputQueueReadCooldown = 50

func processOutputQueue() {
	for {
		if state.OutputQueue.IsEmpty() {
			utils.Sleep(OutputQueueReadCooldown)
			continue
		}
		ctx, e := state.OutputQueue.Dequeue()
		if e != nil {
			utils.Log(e.Error())
			continue
		}
		go sendMessageContextToInner(ctx)
		go sendMessageContextToTargetConnections(ctx)
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

func sendMessageContextToTargetConnections(ctx *MessageContext) {
	if len(ctx.TargetConnections) == 0 {
		return
	}
	for _, targetConnection := range ctx.TargetConnections {
		go targetConnection.Send(ctx.Message.Serialize())
	}
}

// Spec:
// First two bytes -> code
// Rest is to the body.
func newMessage(raw []byte) (*Message, *utils.Error) {
	if len(raw) < 2 {
		return nil, utils.NewError(CodeIncorrectMessageStructure)
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
func Publish(code Code, body MessageBody, options *PublishOptions) *utils.Error {
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
