const Db = @import("./Lib/Database.zig");
const Utils = @import("./Lib/Utils.zig");
const Timeline = @import("./Timeline.zig");

pub const StateEvent = struct {
    Id: Db.Id,
    TypeId: Db.Id,
    TimelineDay: Timeline.Day,
    TimeMs: Utils.TimeMs,
    Body: Utils.String,
};
