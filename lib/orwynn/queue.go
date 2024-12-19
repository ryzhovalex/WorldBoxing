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

func (self *MessageQueue) Enqueue(item *MessageContext) *utils.Error {
	if self.IsFull() {
		return utils.NewError(CodeMessageQueueFull)
	}
	self.data = append(self.data, item)
	return nil
}

func (self *MessageQueue) Dequeue() (*MessageContext, *utils.Error) {
	return nil, nil
}

func (self *MessageQueue) Size() int {
	return len(self.data)
}

func (self *MessageQueue) IsFull() bool {
	return self.Size() == self.maxSize
}

func (self *MessageQueue) IsEmpty() bool {
	return self.Size() == 0
}
