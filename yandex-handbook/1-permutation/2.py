# Import math Library
import math

s = input()

n, k = (int(x) for x in s.split())

# Print total number of possible combinations
print(math.comb(k + n - 1, k))
