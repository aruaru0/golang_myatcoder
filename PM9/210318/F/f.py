import decimal as dec

X, Y, R = input().split()

x = dec.Decimal(X)
y = dec.Decimal(Y)
r = dec.Decimal(R)

top = (y - r).__ceil__()
bottom = (y + r).__floor__()


ans = 0
for i in range(top, bottom+1) :
    left = x - (r**2 - (i-y)**2).sqrt()
    right = x + (r**2 - (i-y)**2).sqrt()
    cnt = right.__floor__() - left.__ceil__()
    ans += cnt + 1

print(ans)