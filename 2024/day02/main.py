import itertools as it

with open("input.txt", "r") as f:
    rows = [[int(n) for n in line.split()] for line in f]

def part1(row: list[int]) -> bool:
    diffs = [j - i for i, j in it.pairwise(row)]
    # All d > 0 or all d < 0 for a row to be strictly increasing or decreasing,
    # so len of the set must be 1
    if len({d > 0 for d in diffs}) == 1:
        return all(1 <= abs(d) <= 3 for d in diffs)
    return False

def part2(row: list[int]) -> bool:
    # Does any row satisfy the part1 constraints when we exclude a value?
    return any(part1([*row[:i], *row[i+1:]]) for i in range(len(row)))

print(f"Part 1: {sum(map(part1, rows))}")
print(f"Part 2: {sum(map(part2, rows))}")
