N = int(input())
a = set(map(int, input().split()))


pos = 0
while N >= 0:
    pos+=1
    if pos in a: N -= 1
    else: N-=2

print(pos-1)