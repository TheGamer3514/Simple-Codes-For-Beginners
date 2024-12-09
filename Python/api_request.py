"""
api_request.py
This script demonstrates making a basic HTTP GET request using the requests library.
"""

import requests

response = requests.get("https://api.github.com")
print("Status code:", response.status_code)
print("Headers:", response.headers)