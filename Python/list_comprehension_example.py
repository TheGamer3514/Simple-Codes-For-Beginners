"""
list_comprehension_example.py
This script demonstrates the use of list comprehensions to create a list of squares.
"""

numbers = [1, 2, 3, 4, 5]
squares = [x ** 2 for x in numbers]
print("Original list:", numbers)
print("List of squares:", squares)