import unittest

from trees import Node, tree_by_levels


class TestsProvided(unittest.TestCase):
    def test_0(self):
        self.assertEqual(tree_by_levels(None), [])
        self.assertEqual(tree_by_levels(Node(Node(None, Node(None, None, 4), 2),
                                             Node(Node(None, None, 5),
                                                  Node(None, None, 6),
                                                  3),
                                             1)),
                         [1, 2, 3, 4, 5, 6])


class TestBase(unittest.TestCase):
    def testRoot(self):
        self.assertEqual(tree_by_levels(Node(None, None, 1)),
                         [1])

    def testSimple(self):
        self.assertEqual(tree_by_levels(Node(Node(None, None, 2),
                                             Node(None, None, 3),
                                             1)),
                         [1, 2, 3])
