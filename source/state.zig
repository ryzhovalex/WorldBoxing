const Db = @import("./lib/Database.zig");
const Utils = @import("./lib/Utils.zig");
const Timeline = @import("./Timeline.zig");

pub const StateEvent = struct {
    Id: Db.Id,
    TypeId: Db.Id,
    TimelineDay: Timeline.Day,
    TimeMs: Utils.TimeMs,
    Body: Utils.String,
};
