const std = @import("std");
const Driver = @import("fridge");
pub var Session: Driver.Session = undefined;
pub const Id = u32;

pub fn Init() !void {
    Session = try Driver.Session.open(
        Driver.SQLite3,
        std.heap.page_allocator,
        .{
            .filename = "var/main.db"
        }
    );
    defer Session.deinit();
    try Session.exec("BEGIN", .{});
}

pub fn Commit() !void {
    try Session.exec("COMMIT", .{});
}

pub fn Rollback() !void {
    try Session.exec("ROLLBACK", .{});
}
