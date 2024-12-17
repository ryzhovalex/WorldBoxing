const std = @import("std");
const utils = @import("./utils.zig");
const commands = std.StaticStringMap(*const fn (utils.String) void).initComptime(.{
    .{ "w", write },
    .{ "r", rollback },
    .{ "q", exit },
});

pub fn Execute(command: utils.String) !void {
    const function = commands.get(command);
    if (function == null) {
        return utils.Error.UnrecognizedCommand;
    }
    function.?(command);
}

fn write() void {}

fn rollback() void {}

fn exit(command: utils.String) void {
    _ = command;
    std.process.exit(0);
}
