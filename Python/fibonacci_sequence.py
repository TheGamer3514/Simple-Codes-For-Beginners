"""
fibonacci_sequence.py
This script demonstrates generating the Fibonacci sequence up to a certain number of terms.
"""

terms = 10
a, b = 0, 1
print("Fibonacci sequence:")
for _ in range(terms):
    print(a, end=" ")
    a, b = b, a + b