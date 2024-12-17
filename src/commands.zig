const std = @import("std");
const utils = @import("./utils.zig");
const db = @import("./db.zig");
pub const Error = error{
    UnrecognizedCommand,
    CommandParsing,
};
const commands = std.StaticStringMap(
    *const fn (CommandContext) void
).initComptime(.{
    .{ "c", commit },
    .{ "r", rollback },
    .{ "q", exit },
});
const CommandContext = struct {
    Raw: *const utils.String,
    Target: *const utils.String,
    Args: [32]*const utils.String,
    Kwargs: [32]*const utils.KeyValue,
};

pub fn Execute(command: utils.String) !void {
    const context = try CreateCommandContext(command);
    const function = commands.get(context.Target.*);
    if (function == null) {
        return Error.UnrecognizedCommand;
    }
    function.?(context);
}

pub fn CreateCommandContext(command: utils.String) !CommandContext {
    var context = CommandContext{
        .Raw = &command,
        .Target = undefined,
        .Args = undefined,
        .Kwargs = undefined,
    };
    var argsIndex: u8 = 0;
    var kwargsIndex: u8 = 0;

    var splitIterator = std.mem.splitScalar(u8, command, ' ');

    var firstPart = true;
    while (splitIterator.next()) |part| {
        if (utils.Contains(u8, part, '=')) {
            if (firstPart) {
                // first part cannot be kwarg
                return Error.CommandParsing;
            }
            var kwargSplitIterator = std.mem.splitScalar(u8, part, '=');
            const key = kwargSplitIterator.next() orelse return utils.Error.Default;
            const value = kwargSplitIterator.next() orelse return utils.Error.Default;
            context.Kwargs[kwargsIndex] = &utils.KeyValue{.K=key, .V=value};
            kwargsIndex += 1;
            continue;
        }
        if (firstPart) {
            context.Target = &part;
            firstPart = false;
            continue;
        }
        context.Args[argsIndex] = &part;
        argsIndex += 1;
    }

    return context;
}

fn commit(context: CommandContext) void {
    _ = context;
}

fn rollback(context: CommandContext) void {
    _ = context;
}

fn exit(_: CommandContext) void {
    std.process.exit(0);
}
