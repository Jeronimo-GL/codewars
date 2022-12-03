import unittest
from secrets import recoverSecret


class TestBasic(unittest.TestCase):
    def test_empty(self):
        secret = []
        self.assertEqual(recoverSecret(secret), "")

    def test_single(self):
        secret = [['a', 'b', 'c']]
        self.assertEqual(recoverSecret(secret), 'abc')

    def test_provided(self):
        secret = "whatisup"
        triplets = [
            ['t', 'u', 'p'],
            ['w', 'h', 'i'],
            ['t', 's', 'u'],
            ['a', 't', 's'],
            ['h', 'a', 'p'],
            ['t', 'i', 's'],
            ['w', 'h', 's']
            ]
        self.assertEqual(recoverSecret(triplets), secret)
