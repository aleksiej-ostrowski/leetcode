import unittest
import _551 as mycode

class Test_All(unittest.TestCase):

    def test_1(self):
        self.assertEqual(mycode.Solution().checkRecord(s="PPALLP"), True, "Must be True")

    def test_2(self):
        self.assertEqual(mycode.Solution().checkRecord(s="PPALLL"), False, "Must be False")

if __name__ == '__main__':
    unittest.main()
