import unittest
from main import get_generation

class GameOfLifeTest(unittest.TestCase):
	def __init__(self, methodName: str = "runTest") -> None:
		self.emptyCase = [[0,0,0], [0,0,0], [0,0,0]]
		super().__init__(methodName)
		
	def testEmpty(self):
		self.assertEquals(get_generation([[]],1), [[]])

	def testZeros(self):
		zeros = [[[0]],
	   		[[0,0], [0,0]],
			[[0,0], [0,0]],
			[[0,0,0], [0,0,0], [0,0,0]],
	   	]
		for z in zeros:
			self.assertEquals(get_generation(z,1), [[]])
			
	def testOneDies(self):
		ones = [
			[[1]],
			[[1,0], [0,0]],
			[[0,1], [0,0]],
			[[0,0], [1,0]],
			[[0,0], [0,1]],
			[[1,0,0], [0,0,0], [0,0,0]],
			[[0,1,0], [0,0,0], [0,0,0]],
			[[0,0,1], [0,0,0], [0,0,0]],
			[[0,0,0], [1,0,0], [0,0,0]],
			[[0,0,0], [0,1,0], [0,0,0]],
			[[0,0,0], [0,0,1], [0,0,0]],
			[[0,0,0], [0,0,0], [1,0,0]],
			[[0,0,0], [0,0,0], [0,1,0]],
			[[0,0,0], [0,0,0], [0,0,1]],
		]
		for o in ones:
			self.assertEquals(get_generation(o,1), [[]])