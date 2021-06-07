x = 10**9
prev = -10

for i in range(x):
    y = x//(i+1)
    if abs(y-prev)<=1 :
        break
    print(i, y)
    prev = y