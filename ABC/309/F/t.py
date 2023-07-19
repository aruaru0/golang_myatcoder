from collections import defaultdict

N = int(input())
hako_list = [list(map(int, input().split())) for _ in range(N)]

print(hako_list)

d=defaultdict(list)
for x,y,z in hako_list:
	d[x].append((y,z))
for x,yz_list in sorted(d.items()):
	for y,z in yz_list :
		if seg.prod(0,y)<z:
			print("Yes")
			exit()
	for y,z in yz_list :
		if z<seg.get(y):
			seg.set(y,z)
			
print("No")