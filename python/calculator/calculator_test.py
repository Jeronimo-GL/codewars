import unittest

from calculator import calculate

class TestCalculator(unittest.TestCase):
    def test_orginales(self):
        cases = [
            ("1.1", 1.1),                   # returns the number if no commands given
            ("10+5", 15),                   # addition
            ("8-2", 6),                     # subtraction
            ("4*3", 12),                    # muliplication
            ("18$2", 9),                    # division
            ("5+8-8*2$4", 9),               # multiple commands
            ("3x+1", "400: Bad request")    # handles incorrect input
            ]
        for (case, result) in cases:
            self.assertEqual(calculate(case), result) 
