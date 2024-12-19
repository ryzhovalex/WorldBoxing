// Making reactivity.
package orwynn

import "worldboxing/lib/utils"

type Code = utils.Code
type Message struct {
	Code Code `json:"Code"`
	Body any  `json:"Body"`
}
type MessageContext struct {
	Message      *Message
	ConnectionId utils.Id
}
type Subscriber func(*Message)
type Subscription struct {
	Id         utils.Id
	Code       Code
	Subscriber Subscriber
}

const MessageQueueMaxSize = 65536

type State struct {
	queue               *MessageQueue
	transport           *Transport
	codeToSubscriptions map[Code][]Subscription
	// TODO: Eliminate risk of available subscription id overflow while
	//		 earlier slots became vacant (by unsubscribing).
	nextSubscriptionId utils.Id
}

type PublishOptions struct {
	TargetConnectionIds []utils.Id
}

type UnsubscribeFunction func()

var state *State

func Init(transport *Transport) *utils.Error {
	state = &State{
		queue:     NewMessageQueue(MessageQueueMaxSize),
		transport: transport,
	}
	return nil
}

func Publish(code Code, body any, options *PublishOptions) *utils.Error {
	return nil
}

func Subscribe(code Code, subscriber Subscriber) (func(), *utils.Error) {
	subscriptionId := state.nextSubscriptionId
	state.nextSubscriptionId += 1
	return func() {
		subscriptions, ok := state.codeToSubscriptions[code]
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
		state.codeToSubscriptions[code] = utils.RemoveFromUnorderedSlice(
			state.codeToSubscriptions[code],
			indexToDelete,
		)
	}, nil
}
