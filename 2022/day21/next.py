def solve(e):
    actions = m[e]
    if len(actions) == 1:
        return actions[0]

    a, op, b = actions
    return "(" + solve(a) + op + solve(b) + ")"


m = {
    monkey: actions.split(" ")
    for line in open(0).read().splitlines()
    for monkey, actions in [line.split(": ")]
}

print(int(eval(solve("root"))))

# Kudos to https://www.reddit.com/r/adventofcode/comments/zrav4h/comment/j133ko6/
m["humn"], m["root"][1] = ["-1j"], "-("
c = eval(solve("root") + ")")
print(int(c.real / c.imag))
