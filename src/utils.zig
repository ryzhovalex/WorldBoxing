pub const Error = error{ Default, UnrecognizedCommand };
pub const String = []const u8;

pub fn TranslateError(e: Error) String {
    switch (e) {
        Error.UnrecognizedCommand => {
            return "Unrecognized command";
        },
        else => {
            return "Error";
        },
    }
}
