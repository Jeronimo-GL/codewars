import re


def alphanumeric(string):
    regex = "[a-zA-Z0-9]+"
    x = re.fullmatch(regex, string)
    return x is not None


if __name__ == "__main__":
    print(alphanumeric("Hello_world"))
