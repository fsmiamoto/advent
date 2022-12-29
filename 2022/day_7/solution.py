import sys
import re
from typing import Dict, List
from functools import reduce

input_lines: List[str] = []
for line in sys.stdin:
    input_lines.append(line.rstrip("\n"))

def deep_get(dictionary, keys, default=None):
    return reduce(lambda d, key: d.get(key, default) if isinstance(d, dict) else default, filter(lambda s : len(s) != 0, keys.split("/")), dictionary)

def go_back(current_dir: str) -> str:
    if current_dir.count('/') == 0:
        return '/'
    return current_dir[:current_dir.rindex('/')]

def change_dir(current_dir, target_dir) -> str:
    if target_dir == '/':
        return '/'
    elif target_dir == '..':
        return go_back(current_dir)

    return current_dir + '/' + target_dir

filesystem = {}
current_dir = ''

cd_regex = r"\$ cd ([a-zA-Z0-9\/\.]+)"
ls = "$ ls"

i = 0

while i < len(input_lines):
    if input_lines[i] == ls:
        current = deep_get(filesystem, current_dir)
        if current is None:
            raise Exception(f"current is None for dir {current_dir}")

        while i+1 < len(input_lines) and not input_lines[i+1].startswith('$'):
            line = input_lines[i+1].split(' ')
            if line[0] == 'dir':
                current[line[1]] = {}
            else:
                current[line[1]] = int(line[0])
            i += 1
    else:
        # cd command
        matches = re.match(cd_regex, input_lines[i])
        if matches is None:
            raise Exception(f"no match for {input_lines[i]}")
        current_dir = change_dir(current_dir, matches.group(1))
    i += 1

result = 0

directories = {}

def get_size(x: Dict | int):
    global result
    if isinstance(x, int):
        return x
    else:
        total = 0
        for key in x:
            for_x = get_size(x[key])
            total += for_x
            if not isinstance(x[key], int):
                directories[key] = for_x
        return total

want = 30000000 - (70000000 - get_size(filesystem))

choosen = None
min_diff = None
for dir in directories:
    if directories[dir] < want:
        continue
    diff = directories[dir] - want
    if not min_diff or diff < min_diff:
        choosen = dir
        min_diff = diff

print(directories[choosen])
