"""
file_write_read.py
This script demonstrates how to write to and read from a file.
"""

filename = "example.txt"
# Write to a file
with open(filename, "w") as file:
    file.write("Hello, File!")

# Read from the file
with open(filename, "r") as file:
    content = file.read()
    print("File content:", content)