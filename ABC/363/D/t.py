N = int(input())
if N == 1:
    print(0)
    exit()

N -= 2
L = 36

for l in range(1, L) :
    h = (l-1)//2
    num = 9*10**h
    if N >= num :
        N -= num
    else :
        print(f"N={N}, h={h}")
        s = str(N + 10**h)
        if l%2 == 0 :
            print(s + s[::-1])
        else :
            print(s + s[-2::-1])
        exit()