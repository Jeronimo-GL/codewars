
digits = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
ops = {'^', '/', '*', '+', '-'}
precedence = {'(': 4, ')': 4, '^': 3, '/': 2, '*': 2, '+': 1, '-': 1}


def to_postfix(infix):
    stack = []
    retstr = ''
    for i in infix:
        if i in digits:
            retstr += i
        if i in ops:
            while len(stack) > 0 \
              and (precedence[stack[0]] >= precedence[i] and not i == '^') \
              and not stack[0] == '(':
                retstr += stack.pop(0)
            stack.insert(0, i)
        elif i == '(':
            stack.insert(0, i)
        elif i == ')':
            while not stack[0] == '(':
                retstr += stack.pop(0)
            if stack[0] == '(':
                stack.pop(0)
    while len(stack) > 0:
        retstr += stack.pop(0)
    return retstr
