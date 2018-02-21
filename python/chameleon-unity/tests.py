import unittest
from  chameleon import chameleon

class TestDefined(unittest.TestCase):
    
    def test_allwrong(self):
        """ There is no solution. -1 expected"""
        self.assertEqual(chameleon([0, 0, 17], 1), -1, msg="All wrong")

    def test_oneStep(self):
        """ Solution in one step"""
        self.assertEqual(chameleon([1, 1, 15], 2), 1, msg="Should return 1")

    def test_averageSituation(self):
        """ Average case """
        self.assertEqual(chameleon([34, 32, 35], 0), 35)

class TestOtherCases(unittest.TestCase):
    
    def test_two_oddd(self):
        """ Two odd numbers to other"""
        self.assertEqual(chameleon([0,1,3], 0), 2)

    def test_odd_and_even(self):
        """ Even and odd numbers to other"""
        self.assertEqual(chameleon([0,1,3], 1), 3)

class TestCornerCases(unittest.TestCase):

    def test_all_zeros(self):
        """ All are zeros"""
        self.assertEqual(chameleon([0,0,0], 1), 0)

    def test_just_one(self):
        """ Just one chameleon """
        self.assertEqual(chameleon([1,0,0], 1), -1)

    def test_just_one_there(self):
        """ One chameleon already there """
        self.assertEqual(chameleon([1,0,0], 0), 0)

    def test_two_not_there(self):
        """ Two chameleon not there """
        self.assertEqual(chameleon([1,1,0], 2), 1)

    def test_two_one_there(self):
        """ Two chameleons one already there """
        self.assertEqual(chameleon([1,1,0], 1), 2)
