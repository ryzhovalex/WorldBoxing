package orwynn

import "worldboxing/lib/utils"

const CodeMessageQueueFull = 2
const CodeMessageQueueEmpty = 3

type MessageQueue struct {
	maxSize int
	data    []*MessageContext
	lock    bool
}

func NewMessageQueue(maxSize int) *MessageQueue {
	return &MessageQueue{
		maxSize,
		[]*MessageContext{},
		false,
	}
}

func (queue *MessageQueue) AcquireLock() {
	for !queue.lock {
		utils.Sleep(10)
	}
	queue.lock = true
}

func (queue *MessageQueue) ReleaseLock() {
	queue.lock = false
}

func (queue *MessageQueue) Enqueue(item *MessageContext) *utils.Error {
	queue.AcquireLock()
	defer queue.ReleaseLock()

	if queue.IsFull() {
		return utils.NewError(CodeMessageQueueFull)
	}
	queue.data = append(queue.data, item)
	return nil
}

// We dequeue all available messages to not mess with lock by single
// operations.
func (queue *MessageQueue) DequeueAll() ([]*MessageContext, *utils.Error) {
	queue.AcquireLock()
	defer queue.ReleaseLock()
	if queue.IsEmpty() {
		return nil, utils.NewError(CodeMessageQueueFull)
	}
	var result []*MessageContext
	copy(result, queue.data)
	queue.data = []*MessageContext{}
	return result, nil
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
