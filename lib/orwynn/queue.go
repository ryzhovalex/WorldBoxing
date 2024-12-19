package orwynn

import "worldboxing/lib/utils"

const CodeMessageQueueFull = 2
const CodeMessageQueueEmpty = 3

type MessageQueue struct {
	maxSize int
	data    []*MessageContext
}

func NewMessageQueue(maxSize int) *MessageQueue {
	return &MessageQueue{
		maxSize,
		[]*MessageContext{},
	}
}

func (queue *MessageQueue) Enqueue(item *MessageContext) *utils.Error {
	if queue.IsFull() {
		return utils.NewError(CodeMessageQueueFull)
	}
	queue.data = append(queue.data, item)
	return nil
}

func (queue *MessageQueue) Dequeue() (*MessageContext, *utils.Error) {
	return nil, nil
}

func (queue *MessageQueue) Size() int {
	return len(queue.data)
}

func (queue *MessageQueue) IsFull() bool {
	return queue.Size() == queue.maxSize
}

func (queue *MessageQueue) IsEmpty() bool {
	return queue.Size() == 0
}
