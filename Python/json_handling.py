"""
json_handling.py
This script demonstrates reading and writing JSON data.
"""

import json

data = {"name": "Alice", "age": 25, "city": "London"}
# Write JSON to a file
with open("data.json", "w") as json_file:
    json.dump(data, json_file)

# Read JSON from the file
with open("data.json", "r") as json_file:
    loaded_data = json.load(json_file)
    print("Loaded JSON:", loaded_data)