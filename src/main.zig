const std = @import("std");
const rl = @import("raylib");

pub fn main() !void {
    const screenWidth = 800;
    const screenHeight = 450;
    rl.initWindow(screenWidth, screenHeight, "World Boxing");
    defer rl.closeWindow();

    rl.setTargetFPS(30);

    while (!rl.windowShouldClose()) {
        rl.beginDrawing();
        defer rl.endDrawing();
        rl.clearBackground(rl.Color.white);
        rl.drawText("Making boxing worldwide.", 190, 200, 20, rl.Color.light_gray);
    }
}
