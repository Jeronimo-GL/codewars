
def solution(number):
    if number <= 1:
        return 0
    else:
        sum = 0
        for i in range(number):
            if i % 3 == 0 or i % 5 == 0:
                sum += i
        return sum


if __name__ == '__main__':
    solution(10)
