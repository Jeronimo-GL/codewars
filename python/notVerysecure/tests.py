import unittest

from notVerySecure import alphanumeric


class TestDefinded(unittest.TestCase):
    def test_1(self):
        self.assertFalse(alphanumeric("hello_world"))

    def test_2(self):
        self.assertTrue(alphanumeric("Passw0rd"))

    def test_3(self):
        self.assertFalse(alphanumeric("   "))

    def test_4(self):
        self.assertTrue(alphanumeric("a"))
