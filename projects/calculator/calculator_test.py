import unittest
from projects.calculator.calculator import Calculator

class Test(unittest.TestCase):
    def test_sum(self):
        calculator = Calculator()
        self.assertEqual(calculator.add(1,2), 4)

    def test_sum_negative(self):
        Calculator = Calculator()
        self.assertEqual(Calculator.add(-4, -5), -9)

if __name__ == '__main__':
    unittest.main()