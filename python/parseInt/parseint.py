def main():
    print("Hello World!")


words = {
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "ten": 10,
    "eleven": 11,
    "twelve": 12,
    "thirteen": 13,
    "fourteen": 14,
    "fifteen": 15,
    "sixteen": 16,
    "seventeen": 17,
    "eighteen": 18,
    "nineteen": 19,
    "twenty": 20,
    "thirty": 30,
    "forty": 40,
    "fifty": 50,
    "sixty": 60,
    "seventy": 70,
    "eighty": 80,
    "ninety": 90,
    "hundred": 100,
    "thousand": 1000,
    "million": 1000000,
}

def parse_word(word):
    if '-' in word:
        parts = word.split('-')
        return words[parts[0]] + words[parts[1]]
    else:
        return words.get(word, None) 


def parse_int(string):
    parts = string.split(' ')
    parts = [parse_word(part) for part in parts if part != 'and']
    num_parts= [p for p in parts if p is not None]
    print(num_parts)
    return  aggregate(num_parts)

def aggregate(numbers):
    if len(numbers) == 1:
        return numbers[0]
    if len(numbers) == 0:
        return 0
    maxpos = numbers.index(max(numbers)) 
    return aggregate(numbers[:maxpos])*numbers[maxpos] + aggregate(numbers[maxpos + 1:])


if __name__ == "__main__":
    print(f"783919 <-> {parse_int('one hundred')}")
