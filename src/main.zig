const std = @import("std");
const Commands = @import("./Commands.zig");
const Utils = @import("./Lib/Utils.zig");
const Database = @import("./Lib/Database.zig");
const Core = @import("./Core.zig");
const Orwynn = @import("./Lib/Orwynn.zig");

pub fn main() !void {
    try Database.Init();
    try Orwynn.Init();
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

        const e = Commands.Execute(slicedInput);
        if (e != null) {
            try stdout.print("[Error] {s}\n", .{Core.TranslateError(e.?)});
        }
    }
}
