"""
nested_loops.py
This script demonstrates the use of nested loops to create a multiplication table.
"""

for i in range(1, 6):
    for j in range(1, 6):
        print(i * j, end="\t")
    print()