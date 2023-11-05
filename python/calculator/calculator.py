import re

opFuncs = [
        ('$', lambda a, b : a/b),
        ('*', lambda a, b : a*b),
        ('-', lambda a, b : a-b),
        ('+', lambda a, b : a+b)
        ]

operators = [s for (s,_) in opFuncs]

def calculate(expression: str):
    partsStr = re.split(r'(\*|\+|-|\$)', expression)
    try:
        parts = [float(p) if not p in operators else p for p in partsStr]
    except:
        return "400: Bad request"
    if len(parts) == 1:
        return parts[0]
    for oper, func in opFuncs:
        while oper in parts:
            pos = parts.index(oper)
            parts[pos] = func(parts[pos-1], parts[pos+1])
            parts[pos-1] = None
            parts[pos+1] = None
            parts = [p for p in parts if not p is None]
    return parts[0]
 

if __name__ == "__main__":
    expression = "5-24*10+20$31-18$48*22+8"
    print(f"{expression} = {calculate(expression)}")
