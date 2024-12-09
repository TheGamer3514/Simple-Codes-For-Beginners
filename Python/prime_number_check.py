"""
prime_number_check.py
This script checks if a given number is a prime number.
"""

def is_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True

number = 29
print(f"{number} is prime:", is_prime(number))