import bisect

def main():
    import sys
    input = sys.stdin.read().split()
    idx = 0
    N = int(input[idx])
    idx += 1
    M = int(input[idx])
    idx += 1

    exchanges = []
    for _ in range(M):
        A = int(input[idx])
        idx += 1
        B = int(input[idx])
        idx += 1
        exchanges.append((A, B))

    # エクスチェンジをA_iでソート
    exchanges.sort()
    sorted_A = [a for a, b in exchanges]
    
    # best_exchanges配列を構築
    best_exchanges = []
    for i in range(len(exchanges)):
        if i == 0:
            best_exchanges.append(exchanges[i])
        else:
            prev_a, prev_b = best_exchanges[-1]
            curr_a, curr_b = exchanges[i]
            if curr_b > prev_b:
                best_exchanges.append((curr_a, curr_b))
            else:
                best_exchanges.append(best_exchanges[-1])

    e = 0
    c = N
    sticker = 0

    while True:
        # 持っているコラをすべて飲む
        e += c
        c = 0

        # 交換可能なオプションがあるか確認
        pos = bisect.bisect_right(sorted_A, e)
        if pos == 0:
            break
        else:
            a, b = best_exchanges[pos - 1]
            e -= a
            c += b
            sticker += 1

    print(sticker)

if __name__ == "__main__":
    main()
