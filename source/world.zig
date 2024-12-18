const db = @import("./lib/Database.zig");
const utils = @import("./lib/Utils.zig");
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

pub const Person = struct {
    Id: db.Id,
    TypeId: db.Id,
    Rating: u32,
    Firstname: utils.String,
    Surname: utils.String,
    CityId: db.Id,
    BornDay: timeline.Day,
    TotalMoneyEarned: f32,
};

pub const CreatePerson = struct {
    Firstname: utils.String,
    Surname: utils.String,
    CityId: db.Id
};

pub fn Create(
    data: CreatePerson
) !void {
    db.Session.insert(
        Person,
        .{
            .Firstname = data.Firstname,
            .Surname = data.Surname,
            .CityId = data.CityId,
            .BornDay = try timeline.GetCurrentDay()
        }
    );
}
