const sqlite = @import("sqlite");
var db: *const sqlite.Db = undefined;

pub fn Init() !void {
    db = &try sqlite.Db.init(.{
        .mode = sqlite.Db.Mode{ .File = "./var/main.db" },
        .open_flags = .{
            .write = true,
        },
        .threading_mode = .MultiThread,
    });
}
