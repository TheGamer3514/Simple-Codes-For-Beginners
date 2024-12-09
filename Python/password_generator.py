"""
password_generator.py
This script demonstrates generating a random password.
"""

import random
import string

def generate_password(length=12):
    characters = string.ascii_letters + string.digits + string.punctuation
    return ''.join(random.choice(characters) for _ in range(length))

print("Generated password:", generate_password())