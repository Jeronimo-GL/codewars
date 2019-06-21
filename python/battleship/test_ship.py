import unittest
from battleship import is_ship_correct

class TestDefined(unittest.TestCase):
    def test_subs_in_border(self):
        field = [[1, 0, 0, 0, 1, 0, 0, 0, 0, 1],
                 [0, 0, 1, 0, 0, 0, 0, 0, 0, 0],
                 [0, 1, 1, 0, 0, 0, 0, 0, 0, 0],
                 [0, 0, 0, 0, 0, 1, 0, 1, 0, 0],
                 [0, 1, 1, 0, 1, 0, 0, 1, 1, 0],
                 [0, 0, 1, 0, 0, 0, 0, 0, 0, 1],
                 [1, 0, 0, 0, 0, 0, 1, 0, 0, 0],
                 [0, 0, 1, 0, 0, 1, 1, 0, 0, 0],
                 [0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
                 [1, 0, 0, 1, 0, 0, 1, 0, 0, 1]]
        self.assertEqual(is_ship_correct(field,(0,0)), 1)
        self.assertEqual(is_ship_correct(field,(0,4)), 1)
        self.assertEqual(is_ship_correct(field,(0,9)), 1)
        self.assertEqual(is_ship_correct(field,(6,0)), 1)
        self.assertEqual(is_ship_correct(field,(9,0)), 1)
        self.assertEqual(is_ship_correct(field,(5,9)), 1)
        self.assertEqual(is_ship_correct(field,(9,6)), 1)
        self.assertEqual(is_ship_correct(field,(9,9)), 1)


    def test_failed_subs(self):
        field = [[1, 0, 0, 0, 1, 0, 0, 0, 0, 1],
                 [0, 0, 1, 0, 0, 0, 0, 0, 0, 0],
                 [0, 1, 1, 0, 0, 0, 0, 0, 0, 0],
                 [0, 0, 0, 0, 0, 1, 0, 1, 0, 0],
                 [0, 1, 1, 0, 1, 0, 0, 1, 1, 0],
                 [0, 0, 1, 0, 0, 0, 0, 0, 0, 1],
                 [1, 0, 0, 0, 0, 0, 1, 0, 0, 0],
                 [0, 0, 1, 0, 0, 1, 1, 0, 0, 0],
                 [0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
                 [1, 0, 0, 1, 0, 0, 1, 0, 0, 1]]
        self.assertEqual(is_ship_correct(field,(9,3)), 0)
        self.assertEqual(is_ship_correct(field,(1,2)), 0)
        self.assertEqual(is_ship_correct(field,(2,1)), 0)
        self.assertEqual(is_ship_correct(field,(2,2)), 0)
        self.assertEqual(is_ship_correct(field,(3,5)), 0)
        self.assertEqual(is_ship_correct(field,(3,7)), 0)
        self.assertEqual(is_ship_correct(field,(9,3)), 0)
        self.assertEqual(is_ship_correct(field,(4,1)), 0)
        self.assertEqual(is_ship_correct(field,(4,2)), 0)
        self.assertEqual(is_ship_correct(field,(4,7)), 0)
        self.assertEqual(is_ship_correct(field,(4,8)), 0)
        self.assertEqual(is_ship_correct(field,(5,2)), 0)
        self.assertEqual(is_ship_correct(field,(6,6)), 0)
        self.assertEqual(is_ship_correct(field,(7,5)), 0)
        self.assertEqual(is_ship_correct(field,(7,6)), 0)
        self.assertEqual(is_ship_correct(field,(8,4)), 0)

