const std = @import("std");
const utils = @import("./utils.zig");
pub const Error = error{
    UnrecognizedCommand,
    CommandParsing,
};
const commands = std.StaticStringMap(*const fn (CommandContext) void).initComptime(.{
    .{ "w", write },
    .{ "r", rollback },
    .{ "q", exit },
});
const CommandContext = struct {
    Raw: utils.String,
    Target: utils.String,
    Args: [32]utils.String,
    Kwargs: [32]utils.KeyValue,
};

pub fn Execute(command: utils.String) !void {
    const context = try CreateCommandContext(command);
    const function = commands.get(context.Target);
    if (function == null) {
        return Error.UnrecognizedCommand;
    }
    function.?(context);
}

pub fn CreateCommandContext(command: utils.String) !CommandContext {
    var target: [32]u8 = undefined;
    var args: [32]utils.String = undefined;
    var kwargs: [32]utils.KeyValue = undefined;
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
            kwargs[kwargsIndex] = utils.KeyValue{.K=key, .V=value};
            kwargsIndex += 1;
            continue;
        }
        if (firstPart) {
            for (part, 0..) |partChar, i| {
                target[i] = partChar;
            }
            firstPart = false;
            continue;
        }
        args[argsIndex] = part;
        argsIndex += 1;
    }

    return CommandContext{
        .Raw = command,
        .Target = &target,
        .Args = args,
        .Kwargs = kwargs,
    };
}

fn write(command: CommandContext) void {
    _ = command;
}

fn rollback(command: CommandContext) void {
    _ = command;
}

fn exit(command: CommandContext) void {
    _ = command;
    std.process.exit(0);
}
