const std = @import("std");
const commands = @import("./commands.zig");
const utils = @import("./lib/utils.zig");
const db = @import("./lib/database.zig");
const core = @import("./core.zig");
const orwynn = @import("./lib/orwynn.zig");

pub fn main() !void {
    try db.Init();
    defer db.Deinit();
    try orwynn.Init();
    defer orwynn.Deinit();

    try cli();
}

fn cli() !void {
    const stdout = std.io.getStdOut().writer();
    const stdin = std.io.getStdIn().reader();
    var buf: [255]u8 = undefined;
    try stdout.print("Welcome to World Boxing!\n", .{});
    while (true) {
        try stdout.print("> ", .{});
        const input = stdin.readUntilDelimiterOrEof(buf[0..], '\n') catch |e| {
            try stdout.print("[Error] {any}\n", .{e});
            continue;
        } orelse {
            break;
        };
        if (input.len == 1 and input[0] == 13) {
            continue;
        }

        // Remove latest `\n`.
        std.mem.reverse(u8, input);
        const slicedInput = input[1..];
        std.mem.reverse(u8, slicedInput);

        const e = commands.Execute(slicedInput);
        if (e != null) {
            try stdout.print("[Error] {s}\n", .{core.TranslateError(e.?)});
        }
    }
}
