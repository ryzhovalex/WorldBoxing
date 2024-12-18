const db = @import("./Database.zig");
const utils = @import("./utils.zig");
const timeline = @import("./timeline.zig");

pub const WorldEvent = struct {
    Id: db.Id,
    Body: utils.String,
    TimelineDay: timeline.Day,
};
