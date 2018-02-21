import unittest
from wfw import order_weight

class TestDefined(unittest.TestCase):
    def test_case_1(self):
        """ First case """
        self.assertEquals(order_weight("103 123 4444 99 2000"), "2000 103 123 4444 99")

    def test_case_2(self):
        """ Same sum -> use string order """
        self.assertEquals(order_weight("2000 10003 1234000 44444444 9999 11 11 22 123"), "11 11 2000 10003 22 123 1234000 44444444 9999")

if __name__ == '__main__':
    unittest.main()
