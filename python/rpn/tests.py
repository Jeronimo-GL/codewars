import unittest
from rpn import to_postfix


class Tests(unittest.TestCase):
    def test_1(self):
        self.assertEqual(to_postfix(''), '')

    def test_simple(self):
        self.assertEqual(to_postfix("2+3"), "23+")

    def test_tres(self):
        self.assertEqual(to_postfix('3-4+5'), '34-5+')

    def test_parentesis(self):
        self.assertEqual(to_postfix("2*(3+4)"), "234+*")

    def test_2(self):
        self.assertEqual(to_postfix("2+7*5"), "275*+")

    def test_3(self):
        self.assertEqual(to_postfix("3*3/(7+1)"), "33*71+/")

    def test_4(self):
        self.assertEqual(to_postfix("5+(6-2)*9+3^(7-1)"), "562-9*+371-^+")

    def test_5(self):
        self.assertEqual(to_postfix("(5-4-1)+9/5/2-7/1/7"), "54-1-95/2/+71/7/-")

    def test_6(self):
        self.assertEqual(to_postfix("1^2^3"), "123^^")
