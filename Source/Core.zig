const Orwynn = @import("./Lib/Orwynn.zig");
const Utils = @import("./Lib/Utils.zig");

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
        Error.UnrecognizedCommand => {
            return "Unrecognized command";
        },
        Error.CommandParsing => {
            return "Command parsing";
        },
        else => {
            return "Error";
        },
    }
}
