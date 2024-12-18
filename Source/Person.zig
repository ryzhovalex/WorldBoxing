const db = @import("./Lib/Database.zig");
const utils = @import("./Lib/Utils.zig");
const timeline = @import("./Lib/Timeline.zig");

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
    });
}
