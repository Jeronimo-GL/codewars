import unittest

from roman_numerals import RomanNumerals


class TestBasicos(unittest.TestCase):
    def test_from(self):
        self.assertEqual(RomanNumerals.from_roman('XXI'), 21)
        self.assertEqual(RomanNumerals.from_roman('III'), 3)
        self.assertEqual(RomanNumerals.from_roman('IV'), 4)
        self.assertEqual(RomanNumerals.from_roman('MCM'), 1900)
        self.assertEqual(RomanNumerals.from_roman('MMVIII'), 2008)
        self.assertEqual(RomanNumerals.from_roman('IL'), 49)
        self.assertEqual(RomanNumerals.from_roman('DII'), 502)

    def test_to_roman(self):
        self.assertEqual(RomanNumerals.to_roman(1), 'I')
        self.assertEqual(RomanNumerals.to_roman(3), 'III')
        self.assertEqual(RomanNumerals.to_roman(4), 'IV')
        self.assertEqual(RomanNumerals.to_roman(6), 'VI')
        self.assertEqual(RomanNumerals.to_roman(8), 'VIII')
        self.assertEqual(RomanNumerals.to_roman(9), 'IX')
        self.assertEqual(RomanNumerals.to_roman(14), 'XIV')
        self.assertEqual(RomanNumerals.to_roman(17), 'XVII')
        self.assertEqual(RomanNumerals.to_roman(19), 'XIX')
        self.assertEqual(RomanNumerals.to_roman(20), 'XX')
        self.assertEqual(RomanNumerals.to_roman(22), 'XXII')
        self.assertEqual(RomanNumerals.to_roman(40), 'XL')
        self.assertEqual(RomanNumerals.to_roman(490), 'XD')
        self.assertEqual(RomanNumerals.to_roman(786), 'DCCLXXXVI')
        self.assertEqual(RomanNumerals.to_roman(900), 'CM')
        self.assertEqual(RomanNumerals.to_roman(904), 'CMIV')
        self.assertEqual(RomanNumerals.to_roman(1000), 'M')
        self.assertEqual(RomanNumerals.to_roman(1990), 'MCMXC')
        self.assertEqual(RomanNumerals.to_roman(3888), 'MMMDCCCLXXXVIII')
