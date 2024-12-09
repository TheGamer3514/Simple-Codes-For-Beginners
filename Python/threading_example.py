"""
threading_example.py
This script demonstrates basic multithreading in Python.
"""

import threading

def print_numbers():
    for i in range(5):
        print(i)

thread = threading.Thread(target=print_numbers)
thread.start()
thread.join()
print("Thread finished!")