"""
potato.py
This script demonstrates file handling in Python by writing a sequence of strings 
to a file named 'temp.txt' and then reading and printing its contents.
"""

f = open("temp.txt","w")
for n in range(1,4): 
  f.write(str(n)+" potato\n")
f.write("4\n")
for n in range(5,8):
  f.write(str(n) +" potato\n")
f.write("More!")
f.close()
f = open("temp","r")
print(f.read())
f.close
