N,T=map(int,input().split())
A=list(map(lambda x:int(x)%T,input().split()))
A.sort()
ans=(A[-1]-A[0]+1)//2
print(A, ans)
for n in range(N-1):
  a=A.pop()
  print(a, A[-1])
  ans=min(ans,(A[-1]-(a-T)+1)//2)
print(ans)