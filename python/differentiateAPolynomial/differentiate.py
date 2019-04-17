import re


def differentiate(equation, point):
    d = derivate(parse_polynomial(equation))
    return evaluate(d, point)


def evaluate(poly, point):
    value = 0
    for p in poly:
        value += p[0]*(point ** p[1])
    return value


def derivate(poly):
    derived = []
    for p in poly:
        if p[1] >= 1:
            derived.append((p[0]*p[1], p[1]-1))
    return derived


def parse_polynomial(poly):
    exp = "[-|\+]?[\d]*x?[\^]?[\d*]?"
    itr = re.finditer(exp, poly)
    parsed_poly = []
    for p in itr:
        if p.group(0) != '':
            parsed_poly.append(parse_part(p.group(0)))
    return parsed_poly


def parse_part(poly):
    exp = "([-|\+]?[\d]*)(x?)[\^]?([\d*]?)"
    m = re.match(exp, poly)
    return (parse_coef(m.group(1)), parse_exp(m.group(3), m.group(2)))


def parse_exp(part, x):
    if part == '' and x == 'x':
        return 1
    elif part == '' and x == '':
        return 0
    else:
        return int(part)


def parse_coef(part):
    if part == '' or part == '+':
        return 1
    elif part == '-':
        return -1
    else:
        return int(part)


if __name__ == '__main__':
    # print(parse_part("x^2"), "x^2")
    # print(parse_part("-x^2"), "-x^2")
    # print(parse_part("+x^2"), "+x^2")
    # print(parse_part("-3x^2"), "-3x^2")
    # print(parse_part("2"), "2")
    # print(parse_part("x"), "x")
    # print(parse_part("3x^2"), "3x^2")
    # print(parse_part("x^2"), "x^2")
    # print(parse_part("3x"), "3x")
    print(parse_polynomial("-4^3+5x^2-3x+3"))
