const Orwynn = @import("./Lib/Orwynn.zig");
const core = @import("./core.zig");
const Utils = @import("./utils.zig");

pub const Codes = enum(Orwynn.MessageCode) {
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

pub fn TranslateError(e: anyerror) Utils.String {
    switch (e) {
        core.Error.UnrecognizedCommand => {
            return "Unrecognized command";
        },
        core.Error.CommandParsing => {
            return "Command parsing";
        },
        else => {
            return "Error";
        },
    }
}
