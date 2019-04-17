import unittest

from differentiate import differentiate


class TestDefinded(unittest.TestCase):
    def test_1(self):
        self.assertEqual(differentiate("12x+2", 3), 12)

    def test_2(self):
        self.assertEqual(differentiate("x^2-x", 3), 5)

    def test_3(self):
        self.assertEqual(differentiate("-5x^2+10x+4", 3), -20)

    def test_x(self):
        self.assertEqual(differentiate("x", 3), 1)
        self.assertEqual(differentiate("3x", 3), 3)

    def test_0(self):
        self.assertEqual(differentiate("1", 3), 0)
        self.assertEqual(differentiate("1", 1), 0)
        self.assertEqual(differentiate("34", -1), 0)
