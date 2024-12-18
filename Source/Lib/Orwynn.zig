const State = struct {
    MessageQueue: [65536]Message,
};
var state: *State = undefined;
pub const MessageCode = i16;

pub const Message = struct {
    Code: MessageCode,
    Body: []i32,
};
