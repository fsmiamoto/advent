from typing import List

lines: List[str] = []

with open("./input.txt") as f:
    lines = [l.rstrip("\n") for l in f]


count = 0

for l in lines:
    ranges = [ [int(x) for x in pair.split("-")] for pair in l.split(",")]

    (lower,upper) = (ranges[0],ranges[1]) if ranges[0][0] < ranges[1][0] else (ranges[1],ranges[0])

    if upper[0] <= lower[1]:
        count += 1

print(count)
