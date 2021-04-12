R, X, Y = map(int, input().split())

d = X*X + Y*Y
r = R*R

if r > d :
    print(2)
    exit(0)

c = 1
while 1:
    if r*c*c >= d :
        print(c)
        break
    c+=1