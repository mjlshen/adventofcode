with open("input.txt", "r") as f:
    left, right = zip(*[[int(n) for n in line.split()] for line in f])

# part1 returns the sum of the differences of the two lists after
# they have been sorted
def part1(left: list[int], right: list[int]) -> int:
    return sum([abs(l - r) for (l, r) in zip(sorted(left), sorted(right))])

# part2 returns the sum of each element in left * the number of times
# it occurs in right
def part2(left: list[int], right: list[int]) -> int:
    return sum(l * right.count(l) for l in left)

print(f"Part 1: {part1(left, right)}")
print(f"Part 2: {part2(left, right)}")
