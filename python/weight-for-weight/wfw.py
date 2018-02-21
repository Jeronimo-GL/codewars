
def order_weight(strng):
    # your code
    numbers = strng.split()
    ordered=sorted(numbers, cmp=lambda x,y : comp(x,y))
    return ' '.join(ordered)

def comp(a,b):
    c=cmp(sumDigits(a), sumDigits(b))
    if c != 0 :
        return c
    else:
        return cmp(a,b)

def sumDigits(num):
    s=0
    for c in num:
        s = s + int(c)
    return s

if __name__ == "__main__":
    print order_weight("400 20 30")
