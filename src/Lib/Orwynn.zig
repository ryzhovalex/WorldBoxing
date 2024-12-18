const std = @import("std");

const QueueMaxSize = 65536;
const MessageQueue = struct {
    const Self = @This();
    const Node = struct {
        value: *const Message,
        next: *const Node,
    };

    front: *const Node,
    back: *const Node,
    currentSize: i32,
    maxSize: i32,

    pub fn New() Self {
        return MessageQueue{
            .front = undefined,
            .back = undefined,
            .currentSize = 0,
            .maxSize = QueueMaxSize,
        };
    }

    pub fn Enqueue(self: *Self, message: *const Message) !void {
        if (self.IsFull()) {
            return Error.MessageQueueFull;
        }
        const newNode = Node{
            .value = message,
            .next = undefined,
        };
        if (self.front == undefined and self.back == undefined) {
            self.front = newNode;
            self.back = newNode;
            return;
        }
        self.back.next = &newNode;
    }

    pub fn Dequeue(self: *Self) !*const Message {
        if (self.IsEmpty()) {
            return Error.MessageQueueEmpty;
        }
        const value = self.front.value;
        if (self.front.next != undefined) {
            self.front = self.front.next;
        } else {
            self.front = undefined;
        }
        return value;
    }

    pub fn GetCurrentSize(self: *Self) i32 {
        return self.currentSize;
    }

    pub fn IsFull(self: *Self) bool {
        return self.currentSize == self.maxSize;
    }

    pub fn IsEmpty(self: *Self) bool {
        return self.currentSize == 0;
    }
};
const State = struct {
    MessageQueue: MessageQueue,
};
var state: *const State = undefined;
pub const MessageCode = i16;
pub const Error = error{
    MessageQueueFull,
    MessageQueueEmpty,
};

pub const Message = struct {
    Code: MessageCode,
    Body: []i32,
};

pub const Subscriber = *const fn (*Message) ?anyerror;

pub fn Init() !void {
    state = &State{
        .MessageQueue = MessageQueue.New(),
    };
}

pub fn Publish(code: MessageCode, body: []i32) !void {
    _ = code;
    _ = body;
}

pub fn Subscribe(code: MessageCode) !void {
    _ = code;
}

test "Init" {
    try Init();
    try std.testing.expect(state.*.MessageQueue.backIndex == 0);
    try std.testing.expect(state.*.MessageQueue.frontIndex == -1);
}
