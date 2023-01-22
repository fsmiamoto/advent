from sys import stdin
from typing import List, Tuple

input_lines: List[str] = []
for line in stdin:
    input_lines.append(line)

class Movement:
    def __init__(self, direction: str, steps: int):
        self.velocity = self.__velocity_for_direction(direction)
        self.steps = steps
        self.direction = direction

    def __velocity_for_direction(self, direction: str) -> Tuple[int,int]:
        return { 'L': (-1,0), 'R': (1,0), 'U': (0,1), 'D': (0,-1)}[direction]

    def __str__(self) -> str:
        return f"({self.direction} {self.steps})"

    def __repr__(self) -> str:
        return self.__str__()


movements: List[Movement] = []
for line in input_lines:
    [direction, steps_as_str] = line.split(' ')
    movements.append(Movement(direction,int(steps_as_str)))

head, tail = (0,0), (0,0)

visited_positions = set()
visited_positions.add(tail)

def move(position: Tuple[int,int], velocity: Tuple[int,int]) -> Tuple[int,int]:
    return (position[0]+velocity[0],position[1]+velocity[1])

def distance(a: Tuple[int,int], b: Tuple[int,int]):
    return max(abs(a[0]-b[0]),abs(a[1]-b[1]))

for movement in movements:
    for _ in range(movement.steps):
        previous_head = head
        head = move(head, movement.velocity)
        if distance(head,tail) >= 2:
            tail = previous_head
        visited_positions.add(tail)

print(len(visited_positions))
