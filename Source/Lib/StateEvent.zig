const Db = @import("./Database.zig");
const Utils = @import("./Utils.zig");
const Timeline = @import("./timeline.zig");

pub const WorldEvent = struct {
    Id: Db.Id,
    TypeId: Db.Id,
    TimelineDay: Timeline.Day,
    TimeMs: Utils.TimeMs,
    Body: Utils.String,
};
