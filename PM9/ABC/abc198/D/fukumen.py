#!/usr/bin/python3
import sys

N = 128
s = "0123456789ABCDEFGHIJKLMNOPQRTSUVWXYZ"
word = [['' for i in range(N)] for j in range(128)]
digit = [0 for i in range(len(s))]
l = [0 for i in range(len(s))]
ok = [False for i in range(10)]
ii = 0
jj = 0
carry = 0
solution = 0


def found():
    global solution
    solution += 1
    # print("\n解 %d" % solution)
    for i in range(imax):
        # if (i == imax-1):
        #     print("-"*jmax)
        for j in range(jmax):
            k = jmax-1-j
            c = word[i][k]
            if (c == ''):
                print(" ", end="")
            else:
                print("%d" % digit[s.index(c)], end='')
        print("")
    exit(0)

def tr(sum):
    global ii, jj
    w = word[ii][jj]
    c = 0 if w == '' else s.index(w)
    if (ii < imax-1):
        ii += 1
        d = digit[c]
        if (d < 0):
            d = l[c]
            while(d <= 9):
                if (ok[d]):
                    digit[c] = d
                    ok[d] = False
                    tr(sum+d)
                    ok[d] = True
                d += 1
            digit[c] = -1
        else:
            tr(sum+d)
        ii -= 1
    else:
        jj += 1
        ii = 0
        carry, d = divmod(sum, 10)
        if (digit[c] == d):
            if (jj < jmax):
                tr(carry)
            elif (carry == 0):
                found()
        else:
            if (digit[c] < 0 and ok[d] and d >= l[c]):
                digit[c] = d
                ok[d] = False
                if (jj < jmax):
                    tr(carry)
                elif (carry == 0):
                    found()
                digit[c] = -1
                ok[d] = True
        jj -= 1
        ii = imax-1


# argv = sys.argv
# argv.pop(0)

argv = [input() for _ in range(3)]
imax = len(argv)
jmax = max(map(len, argv))

for i in range(imax):
    argv[i] = argv[i].upper()
    l[s.index(argv[i][0])] = 1
    a = argv[i][-1::-1]
    for j in range(len(a)):
        word[i][j] = a[j]
        c = word[i][j]
        if (c.isalpha()):
            digit[s.index(c)] = -1
        elif (c.isdigit()):
            digit[s.index(c)] = int(c)
        else:
            print("Invalid parameter.")
            exit(1)

for i in range(10):
    ok[i] = True
tr(0)
if (solution == 0):
    print("UNSOLVABLE")
exit(0)
