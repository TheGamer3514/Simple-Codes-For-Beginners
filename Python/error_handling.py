"""
error_handling.py
This script demonstrates basic error handling using try-except blocks.
"""

try:
    num = int(input("Enter a number: "))
    print("You entered:", num)
except ValueError:
    print("That's not a valid number!")