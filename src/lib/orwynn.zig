const std = @import("std");

const MessageQueueMaxSize = 65536;
const MessageQueue = struct {
    const Self = @This();
    const Node = struct {
        value: *const Message,
        next: ?*Node,
    };

    front: ?*Node = null,
    back: ?*Node = null,
    currentSize: i32 = 0,
    maxSize: i32,

    pub fn New(allocator: std.mem.Allocator) *Self {
        var queue = try allocator.create(MessageQueue);
        queue.maxSize = MessageQueueMaxSize;
        return queue;
    }

    pub fn Enqueue(self: *Self, message: *const Message) !void {
        if (self.IsFull()) {
            return Error.MessageQueueFull;
        }
        var newNode = Node{
            .value = message,
            .next = null,
        };
        if (self.front == null and self.back == null) {
            self.front = &newNode;
            self.back = &newNode;
            return;
        }
        self.back.?.next = &newNode;
        self.currentSize += 1;
    }

    pub fn Dequeue(self: *Self) !*const Message {
        if (self.IsEmpty()) {
            return Error.MessageQueueEmpty;
        }
        const value = self.front.value;
        // for the last element of the queue, the front will be set to null
        self.front = self.front.next;
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
    MessageQueue: *MessageQueue,
};
var state: ?*State = null;
pub const MessageCode = i16;
pub const Error = error{
    MessageQueueFull,
    MessageQueueEmpty,
};

pub const Message = struct {
    Code: MessageCode,
    Body: ?[]i32,
};

pub const Subscriber = *const fn (*Message) ?anyerror;

pub fn Init() !void {
    const allocator = std.heap.page_allocator;
    var s = State{
        .MessageQueue = MessageQueue.New(allocator),
    };
    state = &s;
}

pub fn Deinit() void {
    std.heap.page_allocator.destroy(state.?);
    state = null;
}

pub fn Publish(code: MessageCode, body: []i32) !void {
    _ = code;
    _ = body;
}

pub fn Subscribe(code: MessageCode) !void {
    _ = code;
}

test "init" {
    try Init();
    defer Deinit();

    try std.testing.expect(state.?.MessageQueue.front == null);
    try std.testing.expect(state.?.MessageQueue.back == null);
    try std.testing.expect(state.?.MessageQueue.currentSize == 0);
    try std.testing.expect(state.?.MessageQueue.maxSize == MessageQueueMaxSize);
}

test "enqueue message" {
    try Init();
    defer Deinit();

    var queue = state.?.MessageQueue;
    const message = &Message{
        .Code = 0,
        .Body = null,
    };
    try queue.Enqueue(message);
    const front = queue.front orelse unreachable;
    const back = queue.back orelse unreachable;
    try std.testing.expect(front == back);
    try std.testing.expect(front.value == message);
}
