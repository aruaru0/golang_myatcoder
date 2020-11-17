# 拡張ユークリッド互除法
# ax + by = gcd(a, b) となる解 (x, y) の一例と gcd(a, b) を求める
# 引数 (a, b) 返り値 (gcd(a, b), x, y)
def egcd(a, b):
  if a == 0:
    return b, 0, 1
  else:
    g, y, x = egcd(b % a, a)
    return g, x - (b // a) * y, y

def solve(x, k):
  # a(mod-1) + b(-k) = 1
  _, _, b = egcd(mod-1, -k)
  # a(mod-1) + b(-k) = -1 にして、 0 <= b < mod-1 にする

  print(b)
  b *= -1
  b %= mod-1

  print(b)

  n = pow(x, b, mod)
  return n

mod = 10**9+7
t = int(input())
for _ in range(t):
  x, k = map(int, input().split())
  print(solve(x, k))
