import math

n, m, a = [int(x) for x in input().split() ]
print(((m+a-1)//a)*((n+a-1)//a), end="")
