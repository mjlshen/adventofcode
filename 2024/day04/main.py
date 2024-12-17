
class Day04:

    def __init__(self, filename: str):
        with open(filename, "r") as f:
            self.wordsearch = [[x for x in line if x != "\n"] for line in f]
        self.l = len(self.wordsearch)
        self.w = len(self.wordsearch[0])
        self.search_targets = ['XMAS', 'SAMX']

    def __search_row(self) -> int:
        ans = 0
        for row in self.wordsearch:
            for i in range(self.w - 3):
                ans += ''.join(row[i:i+4]) in self.search_targets

        return ans

    def __search_col(self) -> int:
        ans = 0
        for col in zip(*self.wordsearch):
            for i in range(self.l - 3):
                ans += ''.join(col[i:i+4]) in self.search_targets

        return ans

    def __search_diag(self) -> int:
        ans = 0
        for row in range(self.l - 3):
            for col in range(self.w - 3):
                diag_tl_br = ''.join([self.wordsearch[row][col],
                        self.wordsearch[row+1][col+1],
                        self.wordsearch[row+2][col+2],
                        self.wordsearch[row+3][col+3]])
                diag_bl_tr = ''.join([self.wordsearch[row+3][col],
                        self.wordsearch[row+2][col+1],
                        self.wordsearch[row+1][col+2],
                        self.wordsearch[row][col+3]])
                ans += diag_tl_br in self.search_targets
                ans += diag_bl_tr in self.search_targets

        return ans

    def part1(self) -> int:
        return self.__search_row() + self.__search_col() + self.__search_diag()

    def part2(self) -> int:
        ans = 0
        for row in range(1, self.l - 1):
            for col in range(1, self.w - 1):
                if self.wordsearch[row][col] == "A":
                    if (self.wordsearch[row-1][col-1] in ['M', 'S'] and
                        self.wordsearch[row+1][col+1] in ['M', 'S'] and
                        self.wordsearch[row-1][col-1] != self.wordsearch[row+1][col+1]):
                        if (self.wordsearch[row-1][col+1] in ['M', 'S'] and
                        self.wordsearch[row+1][col-1] in ['M', 'S'] and
                        self.wordsearch[row-1][col+1] != self.wordsearch[row+1][col-1]):
                            ans += 1

        return ans

day04 = Day04("input.txt")

print(f"Part 1: {day04.part1()}")
print(f"Part 2: {day04.part2()}")