"""
class_example.py
This script demonstrates defining a class and creating an object in Python.
"""

class Dog:
    def __init__(self, name, breed):
        self.name = name
        self.breed = breed

    def bark(self):
        print(f"{self.name} says Woof!")

my_dog = Dog("Buddy", "Golden Retriever")
my_dog.bark()