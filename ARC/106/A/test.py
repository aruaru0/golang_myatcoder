N = int(input())

for i in range(50):
    for j in range(50):
        n = pow(3,i) + pow(5,j)
        if n == N :
            print(i,j)
            exit()

print(-1)