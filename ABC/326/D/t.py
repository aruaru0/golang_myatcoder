import itertools
N=int(input())
R=input()
C=input()
def head(s):
  for c in s:
    if c!=".":
      return c
for a,b,c in itertools.product(itertools.permutations(range(N)),repeat=3):
  if any(ai==bi or bi==ci or ci==ai for ai,bi,ci in zip(a,b,c)): continue
  print(a,b,c)
  S=[["."]*N for _ in range(N)]
  for i,j in enumerate(a): S[i][j]="A"
  for i,j in enumerate(b): S[i][j]="B"
  for i,j in enumerate(c): S[i][j]="C"
  if "".join(map(head,S))==R and "".join(map(head,zip(*S)))==C:
    print("Yes")
    print(*map(lambda s:"".join(s),S),sep="\n")
    exit()
print("No")
