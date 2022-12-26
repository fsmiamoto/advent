# ugly code ahead
from typing import List
from collections import defaultdict

lines: List[str] = []

with open('./input.txt') as f:
    lines.extend(f.readlines())

def get_score_for_item(item: str):
    if item.islower():
        return ord(item) - 96
    else:
        return ord(item) - 38

grouped = []
for i in range(0,len(lines),3):
    grouped.append([lines[i].rstrip('\n'),lines[i+1].rstrip('\n'),lines[i+2].rstrip('\n')])

total = 0

for group in grouped:
    print(group)
    char_count = defaultdict(int)

    char = None

    in_first = {x for x in group[0]}
    in_second = {x for x in group[1]}

    for c in group[2]:
        if c in in_first and c in in_second:
            char = c

    if char is None:
        raise Exception("Raise your kappas")

    total += get_score_for_item(char)
