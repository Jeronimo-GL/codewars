import unittest

from multiples import solution

class TestsDefined(unittest.TestCase):
    def test_10(self):
        """ test for number 10 """
        self.assertEqual(solution(10), 23)

    def test_20(self):
        """ test for number 20 """
        self.assertEqual(solution(20), 78)


class TestCornerCases(unittest.TestCase):
    def test_0(self):
        """ solution(0) should return 0 """
        self.assertEqual(solution(0), 0)

    def test_1(self):
        """ solution(1) should return 1 """
        self.assertEqual(solution(1), 0)

    def test_big(self):
        """ should work for very big numbers """
        self.assertGreater(solution(10000000), 0)
