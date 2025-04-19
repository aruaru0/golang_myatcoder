def main():
    import sys
    N = int(sys.stdin.readline())
    sets = []
    for _ in range(N):
        parts = list(map(int, sys.stdin.readline().split()))
        C_i = parts[0]
        A_list = parts[1:]
        sets.append(set(A_list))
    
    count = 0
    for mask in range(1, 2**N):
        bits = bin(mask).count('1')
        if bits < 2:
            continue
        selected_sets = []
        for i in range(N):
            if mask & (1 << i):
                selected_sets.append(sets[i])
        common = set.intersection(*selected_sets)
        if not common:
            count += 1
        else:
            valid = True
            for num in common:
                if num % 2 == 0:
                    valid = False
                    break
            if valid:
                count += 1
    print(count)

if __name__ == "__main__":
    main()
