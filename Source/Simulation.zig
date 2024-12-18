const Db = @import("./Lib/Database.zig");
const Utils = @import("./Lib/Utils.zig");
const Timeline = @import("./Timeline.zig");
pub const Error = error{
    WorldAlreadyCreated,
};

pub fn CreateWorld() !void {
    // Timeline shouldn't exists for the world to be considered as not created
    // yet.
    const timelineRow = Db.Session.query(Timeline.Timeline).findFirst() catch {
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
    try Db.Session.insert(Timeline.Timeline, .{
        .CurrentDay = 0,
    });
}
