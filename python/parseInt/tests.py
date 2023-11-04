import unittest

from parseint import parse_int

class TestParseInt(unittest.TestCase):

    def test_single_digit(self):
        self.assertEqual(parse_int("one"), 1)
        self.assertEqual(parse_int("two"), 2)
        self.assertEqual(parse_int("three"), 3)
        self.assertEqual(parse_int("four"), 4)
        self.assertEqual(parse_int("nine"), 9)
        self.assertEqual(parse_int("zero"), 0)

    def test_teens(self):
        self.assertEqual(parse_int("eleven"), 11)
        self.assertEqual(parse_int("twelve"), 12)
        self.assertEqual(parse_int("thirteen"), 13)
        self.assertEqual(parse_int("fourteen"), 14)
        self.assertEqual(parse_int("nineteen"), 19)
        self.assertEqual(parse_int("ten"), 10)


    def test_tens(self):
        self.assertEqual(parse_int("twenty"), 20)
        self.assertEqual(parse_int("thirty"), 30)
        self.assertEqual(parse_int("forty"), 40)
        self.assertEqual(parse_int("fifty"), 50)
        self.assertEqual(parse_int("ninety"), 90)
        self.assertEqual(parse_int("twenty-nine"), 29)


    def test_hundreds(self):
        self.assertEqual(parse_int("one hundred"), 100)
        self.assertEqual(parse_int("two hundred"), 200)
        self.assertEqual(parse_int("three hundred and ten"), 310) 
        self.assertEqual(parse_int("four hundred and twenty"), 420)
        self.assertEqual(parse_int("five hundred ninety-nine"), 599)

    def test_thousands(self):
        self.assertEqual(parse_int("one thousand"), 1000)
        self.assertEqual(parse_int("two thousand"), 2000)
        self.assertEqual(parse_int("three thousand and ten"), 3010) 
        self.assertEqual(parse_int("four thousand and twenty"), 4020)
        self.assertEqual(parse_int("five thousand ninety-nine"), 5099)
        self.assertEqual(parse_int("one thousand and one"), 1001)
        self.assertEqual(parse_int("one thousand three hundred and ten"), 1310) 
        

    def test_parseint(self):
        self.assertEqual(parse_int("one"), 1)
        self.assertEqual(parse_int("twenty"), 20)
        self.assertEqual(parse_int("two hundred forty-six"), 246)
        self.assertEqual(parse_int("seven hundred eighty-three thousand nine hundred and nineteen"), 783919)
        self.assertEqual(parse_int("one million"), 1000000)

