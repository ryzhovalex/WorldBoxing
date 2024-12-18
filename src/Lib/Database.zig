const std = @import("std");
const Sqlite = @import("sqlite");
pub var Session: Sqlite.Db = undefined;
pub const Id = u32;

pub fn Init() !void {
    Session = try Sqlite.Db.init(.{
        .mode = Sqlite.Db.Mode{
            .File="var/main.db"
        },
        .open_flags = .{
            .write = true,
        },
        .threading_mode = .MultiThread,
    });
    defer Session.deinit();
    try Session.exec("BEGIN", .{}, .{});
}

pub fn Commit() !void {
    try Session.exec("COMMIT", .{}, .{});
}

pub fn Rollback() !void {
    try Session.exec("ROLLBACK", .{}, .{});
}
