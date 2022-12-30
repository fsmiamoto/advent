from typing import List, Tuple
from sys import stdin

input_lines: List[str] = []

for line in stdin:
    input_lines.append(line.rstrip("\n"))

tree_map = list(map(lambda s: [int(x) for x in s], input_lines))

def part1(tree_map: List[List[int]]) -> int:
    result = 0
    height, width = len(tree_map), len(tree_map[0])

    def in_border(i,j:int):
        return i == 0 or i == height-1 or j == 0 or j == width-1

    def process(i,j: int) -> bool:
        for (x,y) in [(-1,0),(1,0),(0,1),(0,-1)]:
            if check_direction(i,j,(x,y)):
                return True
        return False

    def check_direction(i,j: int, direction: Tuple[int,int]) -> bool:
        x, y = i,j
        while not in_border(x,y):
            x, y = x + direction[0], y + direction[1]
            if tree_map[x][y] >= tree_map[i][j]:
                return False
        return True


    for i in range(height):
        for j in range(width):
            if process(i,j):
                result += 1

    return result


def part2(tree_map: List[List[int]]) -> int:
    height, width = len(tree_map), len(tree_map[0])

    def in_border(i,j:int):
        return i == 0 or i == height-1 or j == 0 or j == width-1

    def process(i,j: int) -> int:
        result = 1
        for (x,y) in [(-1,0),(1,0),(0,1),(0,-1)]:
                result *= check_direction(i,j,(x,y))
        return result

    def check_direction(i,j: int, direction: Tuple[int,int]) -> int:
        result = 0
        x, y = i,j
        while not in_border(x,y):
            result += 1
            x, y = x + direction[0], y + direction[1]
            if tree_map[x][y] >= tree_map[i][j]:
                break
        return result


    result = None
    for i in range(height):
        for j in range(width):
            if not result or process(i,j) > result:
                result = process(i,j)

    if result is None:
        raise Exception("deu ruim")

    return result


print(part1(tree_map))
print(part2(tree_map))
