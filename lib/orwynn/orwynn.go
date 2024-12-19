// Making reactivity.
package orwynn

import "worldboxing/lib/utils"

const CodeCannotReferenceInnerTransport = 2

type Code = utils.Code
type Message struct {
	Code Code `json:"Code"`
	Body any  `json:"Body"`
}
type MessageContext struct {
	Message *Message
	// AuthorConnection of message's author. Set to nil if author is the host.
	AuthorConnection *Connection
	// Connections to which the message should be sent to. Only considered for
	// output messages.
	TargetConnections []*Connection
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
	Transports          map[string]*Transport
	CodeToSubscriptions map[Code][]*Subscription
	// TODO: Eliminate risk of available subscription id overflow while
	//		 earlier slots became vacant (by unsubscribing).
	NextSubscriptionId utils.Id
}

type PublishOptions struct {
	TargetConnections []*Connection
}

type UnsubscribeFunction func()

var state *State

func Init(transports map[string]*Transport) *utils.Error {
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
	return nil
}

// Publish a message body to the bus.
//
// If options.TargetConnections is nil/empty, the inner transport will be used.
func Publish(code Code, body any, options *PublishOptions) *utils.Error {
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
