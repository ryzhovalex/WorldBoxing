const db = @import("./db.zig");
const utils = @import("./utils.zig");
const timeline = @import("./timeline.zig");

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
