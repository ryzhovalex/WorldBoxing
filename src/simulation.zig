const db = @import("./db.zig");
const utils = @import("./utils.zig");
const timeline = @import("./timeline.zig");
pub const Error = error{
    WorldAlreadyCreated,
};

pub fn CreateWorld() !void {
    // Timeline shouldn't exists for the world to be considered as not created
    // yet.
    const timelineRow = db.Session.query(timeline.Timeline).findFirst() catch {
        return try createWorld();
    };
    if (timelineRow == null) {
        return try createWorld();
    }
    return Error.WorldAlreadyCreated;
}

fn createWorld() !void {
}

fn createTimeline() !void {
    try db.Session.insert(timeline.Timeline, .{
        .CurrentDay = 0,
    });
}
