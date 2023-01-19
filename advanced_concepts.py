# 1. Functions:

# Defining a function
def add(x, y):
    return x + y

# Calling a function
result = add(3, 4)
print(result) # Output: 7

# Functions can also have default values for their arguments:
def add_with_default(x, y=2):
    return x + y
print(add_with_default(5)) # Output: 7

# Functions can also be defined using Lambda expressions
add = lambda x, y : x + y
print(add(5,6)) # Output: 11

# 2. Modules:

# Importing a module
import math

# Using a function from the math module
result = math.sqrt(16)
print(result) # Output: 4.0

# Importing only specific functions from a module
from math import sqrt, pi
result = sqrt(16)
print(result) # Output: 4.0
