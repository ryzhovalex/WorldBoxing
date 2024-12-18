const std = @import("std");
const Utils = @import("./Lib/Utils.zig");
const Database = @import("./Lib/Database.zig");
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
    Raw: *const Utils.String,
    Target: *const Utils.String,
    Args: [32]*const Utils.String,
    Kwargs: [32]*const Utils.KeyValue,
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

pub fn Execute(command: Utils.String) ?anyerror {
    const context = CreateCommandContext(command) catch |e| {
        return e;
    };
    const function = commands.get(context.Target.*) orelse return Core.Error.UnrecognizedCommand;
    const e = function(context);
    return e;
}

pub fn CreateCommandContext(command: Utils.String) !CommandContext {
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
        if (Utils.Contains(u8, part, '=')) {
            if (firstPart) {
                // first part cannot be kwarg
                return Core.Error.CommandParsing;
            }
            var kwargSplitIterator = std.mem.splitScalar(u8, part, '=');
            const key = kwargSplitIterator.next() orelse return Utils.Error.Default;
            const value = kwargSplitIterator.next() orelse return Utils.Error.Default;
            context.Kwargs[kwargsIndex] = &Utils.KeyValue{.K=key, .V=value};
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
