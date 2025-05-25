import sys
import heapq

def main():
    input = sys.stdin.read().split()
    ptr = 0
    T = int(input[ptr])
    ptr += 1
    for _ in range(T):
        N = int(input[ptr])
        ptr += 1
        A = []
        for _ in range(2 * N):
            A.append(int(input[ptr]))
            ptr += 1
        total_sum = sum(A)
        heap = []
        for j in range(2 * N):
            i = j + 1
            val = A[j]
            heapq.heappush(heap, -val)
            if len(heap) > (i // 2):
                heapq.heappop(heap)
        sum_selected = -sum(heap)
        print(total_sum - sum_selected)

if __name__ == "__main__":
    main()
