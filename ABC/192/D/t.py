x = input()
m = int(input())

if len(x) == 1:
    if int(x) <= m:
        print(1)
    else:
        print(0)
    exit()

def f(s, base) :
    ret = 0
    for c in s:
        ret *= base
        ret += int(c)
    return ret

d = int(max(list(x)))

l = d
r = int(1e18)+1

while r - l > 1 :
    mid = (l+r)//2
    if f(x, mid) <= m :
        l = mid
    else :
        r = mid

r -= d + 1
print(r)