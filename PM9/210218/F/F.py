from decimal import *

s = input().split()

x = Decimal(s[0])
y = Decimal(s[1])
r = Decimal(s[2])

top = (x-r).__ceil__()
btm = (x+r).__floor__()

ans = 0
for i in range(top, btm+1):
    left = y - (r**2 - (i-x)**2).sqrt()
    right = y + (r**2 - (i-x)**2).sqrt()
    cnt = right.__floor__() - left.__ceil__()
    ans += cnt +1

print(ans)