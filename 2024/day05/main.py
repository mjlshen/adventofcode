from functools import cmp_to_key

class Day05:
    def __init__(self, filename: str):
        ordering_rules, pages = open(filename).read().split('\n\n')
        # In Python, a comparison function returns:
        # * A negative number if first arg < second arg
        # * Zero if the two args are equal
        # * A positive number if second arg > first arg
        # If x|y exists in ordering_rules, return -True --> -1
        # i.e. if a rule exists for x|y, x should be before y
        self.cmp = cmp_to_key(lambda x, y: -(x+'|'+y in ordering_rules))
        self.updates = pages.split()

    def part1(self) -> int:
        ans = 0
        for u in self.updates:
            u = u.split(',')
            s = sorted(u, key=self.cmp)
            if u == s:
                ans += int(s[len(s) // 2])

        return ans

    def part2(self) -> int:
        ans = 0
        for u in self.updates:
            u = u.split(',')
            s = sorted(u, key=self.cmp)
            if u != s:
                ans += int(s[len(s) // 2])

        return ans

day05 = Day05("input.txt")
print(f"Part 1: {day05.part1()}")
print(f"Part 2: {day05.part2()}")