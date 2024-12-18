const std = @import("std");
const db = @import("fridge");
pub var Session: db.Session = undefined;
pub const Id = u32;

pub fn Init() !void {
    const allocator = std.heap.page_allocator;
    Session = try db.Session.open(
        db.SQLite3, allocator, .{ .filename = "var/main.db" }
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
