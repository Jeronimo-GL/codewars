MAX = 2


def recoverSecret(triplets):
    secret = []
    while len(triplets) > 0:
        print(triplets)
        print(secret)
        t = triplets.pop(0)
        if len(secret) == 0:
            secret = t
        elif t[0] in secret:
            secret = merge_after(t, secret)
        elif t[MAX] in secret:
            secret = merge_before(t, secret)
        else:
            triplets.append(t)

    return "".join(secret)


def merge_before(arr1, arr2):
    '''
    Merges arr1 with arr2 concatenating arr2 into arr1 begining with the
    position of the last char of arr1 in arr2
    '''

    i = arr2.index(arr1[-1])
    if i <= MAX:
        return arr1 + arr2[i+1:]
    else:
        return arr2


def merge_after(arr1, arr2):
    '''
    Merges arr1 with arr2 concatenating arr1 after arr2 begining with the
    position of the first char in arr1 present in arr2
    '''
    i = arr2.index(arr1[0])
    return arr2[:i] + arr1


def main2():
    arr1 = ['d', 'e', 'f']
    arr2 = ['c', 'd']
    print(merge_after(arr1, arr2))


def main1():
    triplets = [
        ['t', 'u', 'p'],
        ['w', 'h', 'i'],
        ['t', 's', 'u'],
        ['a', 't', 's'],
        ['h', 'a', 'p'],
        ['t', 'i', 's'],
        ['w', 'h', 's']
    ]
    print(recoverSecret(triplets))


if __name__ == "__main__":
    main1()
