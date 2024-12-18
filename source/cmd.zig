const std = @import("std");
const utils = @import("./lib/utils.zig");
const Database = @import("./lib/database.zig");
const Core = @import("./Core.zig");
const Simulation = @import("./Simulation.zig");

const commands = std.StaticStringMap(
    *const fn (CommandContext) ?anyerror
).initComptime(.{
    .{ "c", commit },
    .{ "r", rollback },
    .{ "q", exit },

    .{ "simulation.CreateWorld", simulationCreateWorld },
});
const CommandContext = struct {
    Raw: *const utils.String,
    Target: *const utils.String,
    Args: [32]*const utils.String,
    Kwargs: [32]*const utils.KeyValue,
};

fn simulationCreateWorld(_: CommandContext) ?Core.Error {
    Simulation.CreateWorld() catch return Core.Error.Default;
    return null;
}

fn commit(_: CommandContext) ?Core.Error {
    Database.Commit() catch return Core.Error.Default;
    return null;
}

fn rollback(_: CommandContext) ?Core.Error {
    Database.Rollback() catch return Core.Error.Default;
    return null;
}

fn exit(_: CommandContext) ?Core.Error {
    std.process.exit(0);
    return null;
}

pub fn Execute(command: utils.String) ?anyerror {
    const context = CreateCommandContext(command) catch |e| {
        return e;
    };
    const function = commands.get(context.Target.*) orelse return Core.Error.UnrecognizedCommand;
    const e = function(context);
    return e;
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
                return Core.Error.CommandParsing;
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
