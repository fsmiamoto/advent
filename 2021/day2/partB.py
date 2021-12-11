from sys import stdin;

horizontal, aim, depth = 0,0,0

for line in stdin:
    [cmd, value] = line.split()
    if cmd == 'forward':
        horizontal += int(value)
        depth += int(value) * aim
    elif cmd == 'down':
        aim += int(value)
    elif cmd == 'up':
        aim -= int(value)

print(horizontal*depth)
