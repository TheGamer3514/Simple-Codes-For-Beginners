"""
factorial_calculation.py
This script demonstrates calculating the factorial of a number using recursion.
"""

def factorial(n):
    if n == 0:
        return 1
    return n * factorial(n - 1)

num = 5
print(f"Factorial of {num} is {factorial(num)}")