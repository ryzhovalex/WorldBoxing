const orwynn = @import("./Lib/Orwynn.zig");
pub const Codes = enum(orwynn.MessageCode) {
    PersonCreated = -2,
    CreatePerson = -1,

    Ok = 0,

    Error = 1,
};
pub const Error = error{
    Default,
    UnrecognizedCommand,
    CommandParsing,
};
