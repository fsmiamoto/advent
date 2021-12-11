from sys import stdin;

horizontal, depth = 0,0

for line in stdin:
    [cmd, value] = line.split()
    if cmd == 'forward':
        horizontal += int(value)
    elif cmd == 'down':
        depth += int(value)
    elif cmd == 'up':
        depth -= int(value)

print(horizontal*depth)
