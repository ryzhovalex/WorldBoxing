const std = @import("std");

const MessageQueueMaxSize = 65536;
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
            .maxSize = MessageQueueMaxSize,
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
        self.currentSize += 1;
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
        self.currentSize -= 1;
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
var state: *State = undefined;
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
    var s = State{
        .MessageQueue = MessageQueue.New(),
    };
    state = &s;
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
    try std.testing.expect(state.MessageQueue.front == undefined);
    try std.testing.expect(state.MessageQueue.back == undefined);
    try std.testing.expect(state.MessageQueue.currentSize == 0);
    try std.testing.expect(state.MessageQueue.maxSize == MessageQueueMaxSize);
}

test "Enqueue message" {
    try Init();
    state.MessageQueue.Enqueue(&Message{
        .Code = 0,
        .Body = undefined,
    });
    std.testing.expect(state.MessageQueue.front != undefined);
    std.testing.expect(state.MessageQueue.back != undefined);
}
