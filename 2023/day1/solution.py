def get_lines():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()] 

def part1():
    lines = get_lines()

    values = []
    for line in lines:
        digits = [char for char in line if char.isdigit()]
        values.append(int(digits[0] + digits[-1]))
    
    print(sum(values))

def part2():
    pass

if __name__ == "__main__":
    part1()
    part2()
