def anagrams(word, words):
    wl = sorted(list(word))
    retList = []
    print("wl", wl)
    for w in words:
        ll = sorted(list(w))
        print("ll", ll)
        if ll == wl:
            retList.append(w)
    return retList


if __name__ == '__main__':
    print(anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada']))
