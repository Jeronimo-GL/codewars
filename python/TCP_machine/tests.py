import unittest
from tcp import traverse_TCP_states

class TestsProvided(unittest.TestCase):
    def test_0(self):
        self.assertEqual(traverse_TCP_states(["APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN"]), "CLOSE_WAIT")
        self.assertEqual(traverse_TCP_states(["APP_PASSIVE_OPEN",  "RCV_SYN","RCV_ACK"]), "ESTABLISHED")
        self.assertEqual(traverse_TCP_states(["APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN","APP_CLOSE"]), "LAST_ACK")
        self.assertEqual(traverse_TCP_states(["APP_ACTIVE_OPEN"]), "SYN_SENT")
        self.assertEqual(traverse_TCP_states(["APP_PASSIVE_OPEN","RCV_SYN","RCV_ACK","APP_CLOSE","APP_SEND"]), "ERROR")

class TestsFinal(unittest.TestCase):
    def test_1(self):
        self.assertEqual(traverse_TCP_states(["APP_ACTIVE_OPEN","RCV_SYN","APP_CLOSE","RCV_FIN","RCV_ACK","APP_TIMEOUT"]), "CLOSED")
