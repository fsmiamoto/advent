# import pdb
import re
from typing import Tuple

class Node:
    def __init__(self, value: str, next = None):
        self.value = value
        self.next = next

    def __str__(self) -> str:
        if(self is None):
            return ''
        return self.value + '->' + self.next.__str__()
class Command:
    def __init__(self, quantity: int, from_stack: int, to_stack: int):
        self.quantity = quantity
        self.from_stack = from_stack
        self.to_stack = to_stack

# Didn't want to parse that initial input so added it manually :D
stacks = [
    Node('G', Node('J', Node('Z'))),
    Node('C', Node('V', Node('F', Node('W', Node('P', Node('R', Node('L', Node('Q')))))))),
    Node('R', Node('G', Node('L', Node('C', Node('M', Node('P', Node('F'))))))),
    Node('M', Node('H', Node('P', Node('W', Node('B', Node('F', Node('L'))))))),
    Node('Q', Node('V', Node('S', Node('F', Node('C', Node('G')))))),
    Node('L', Node('T', Node('Q', Node('M', Node('Z', Node('J', Node('H', Node('W')))))))),
    Node('V', Node('B', Node('S', Node('F', Node('H'))))),
    Node('S', Node('Z', Node('J', Node('F')))),
    Node('T', Node('B', Node('H', Node('F', Node('P', Node('D', Node('C', Node('M')))))))),
]

commands = []

with open('./input.txt') as f:
    for line in f:
        result = re.search(r"move (\d+) from (\d+) to (\d+)", line)
        commands.append(Command(int(result.group(1)), int(result.group(2)), int(result.group(3))))
def walk(node: Node, num_of_steps: int) -> Node:
    result = node
    while num_of_steps > 1 and result.next:
        result = result.next
        num_of_steps -= 1
    return result

def move(from_stack: Node, to_stack: Node, amount: int) -> Tuple[Node, Node]:
    head = from_stack
    tail = walk(head, amount)

    from_stack = tail.next
    tail.next = to_stack
    to_stack = head
    return from_stack, to_stack

for c in commands:
    from_stack = stacks[c.from_stack-1]
    to_stack = stacks[c.to_stack-1]
    stacks[c.from_stack-1], stacks[c.to_stack-1] = move(from_stack, to_stack, c.quantity)

print(''.join(map(lambda node : node.value, stacks)))
