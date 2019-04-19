import re


def top_3_words(text):
    expr = "[a-zA-Z'*]*[a-zA-Z]+[a-zA-Z'*]*"
    dict = {}
    itr = re.finditer(expr, text)
    for p in itr:
        w = p.group().lower()
        if w not in dict.keys():
            dict[w] = 1
        else:
            dict[w] = dict[w]+1
    s = sorted(dict, key=dict.__getitem__, reverse=True)
    return s[0:3]


if __name__ == '__main__':
    top_3_words("wBSo'DnJ'G:;wBSo'DnJ'G?.wBSo'DnJ'G_-,:?wBSo'DnJ'G :wBSo'DnJ'G!;.:wBSo'DnJ'G; ,,_wBSo'DnJ'G  , wBSo'Dn")
