pub const Error = error{ Default };
pub const String = []const u8;
pub const KeyValue = struct { K: String, V: String };

pub fn Contains(comptime T: type, array: []const T, item: T) bool {
    for (array) |thing| {
        if (thing == item) {
            return true;
        }
    }
    return false;
}
