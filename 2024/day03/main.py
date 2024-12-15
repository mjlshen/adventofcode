import re

data = open("input.txt").read().replace('\n','')

def part1(data: str) -> int:
    return sum(int(x)*int(y) for x,y in re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", data))

def part2(data: str) -> int:
    # Remove all 'don't().*do()' sections from data, including any trailing don't()s
    data = re.sub(r"don't\(\)(.*?)do\(\)", "", data + "do()")
    return part1(data)

print(f"Part 1: {part1(data)}")
print(f"Part 2: {part2(data)}")
