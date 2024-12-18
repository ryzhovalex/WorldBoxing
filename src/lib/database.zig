const std = @import("std");
const driver = @import("fridge");
pub var Session: driver.Session = undefined;
pub const Id = u32;

pub fn Init() !void {
    Session = try driver.Session.open(
        driver.SQLite3,
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
