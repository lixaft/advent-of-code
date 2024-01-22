const std = @import("std");
const input = @embedFile("input.txt");

pub fn main() !void {
    var result: u32 = 0;

    var tokens = std.mem.tokenizeScalar(u8, input, '\n');
    while (tokens.next()) |token| {
        var first: ?u8 = null;
        var last: ?u8 = null;

        for (token) |c| {
            if (c < '0' or c > '9') {
                continue;
            }

            const d = c - '0';
            if (first == null) {
                first = d;
            }
            last = d;
        }

        result += first.? * 10 + last.?;
    }
    std.debug.print("{d}\n", .{result});
}
