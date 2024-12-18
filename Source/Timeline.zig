const db = @import("./Database.zig");
const utils = @import("./utils.zig");
pub const Day = u32;
const currentDay: ?Day = null;

pub const Timeline = struct {
    Id: db.Id,
    CurrentDay: Day,
};

pub fn GetCurrentDay() !Day {
    if (currentDay == null) {
        const timeline = try db.Session.query(Timeline).findFirst() orelse return utils.Error.Default;
        currentDay = timeline.CurrentDay;
    }
    return currentDay;
}
