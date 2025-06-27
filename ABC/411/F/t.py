import sys

def main():
    input = sys.stdin.read().split()
    ptr = 0
    N = int(input[ptr])
    ptr += 1
    M = int(input[ptr])
    ptr += 1

    edges = []
    global_set = set()
    incident = [set() for _ in range(N + 1)]  # 1-based indexing

    for _ in range(M):
        u = int(input[ptr])
        ptr += 1
        v = int(input[ptr])
        ptr += 1
        if u > v:
            u, v = v, u
        edges.append((u, v))
        global_set.add((u, v))
        incident[u].add((u, v))
        incident[v].add((u, v))

    Q = int(input[ptr])
    ptr += 1
    queries = list(map(int, input[ptr:ptr + Q]))
    ptr += Q

    # DSU implementation
    parent = list(range(N + 1))

    def find(u):
        while parent[u] != u:
            parent[u] = parent[parent[u]]
            u = parent[u]
        return u

    def union(u, v):
        pu = find(u)
        pv = find(v)
        if pu != pv:
            parent[pv] = pu

    for X in queries:
        # Get the original edge
        u_orig, v_orig = edges[X - 1]
        u = find(u_orig)
        v = find(v_orig)
        if u == v:
            print(len(global_set))
            continue
        a = min(u, v)
        b = max(u, v)
        if (a, b) not in global_set:
            print(len(global_set))
            continue
        # Proceed with contraction
        global_set.remove((a, b))
        w = u  # choose u as the new root
        # Process incident edges of u
        for edge in list(incident[u]):
            if edge[0] == u and edge[1] != v:
                if edge in global_set:
                    global_set.remove(edge)
                    x = edge[1]
                    new_a = min(w, x)
                    new_b = max(w, x)
                    if (new_a, new_b) not in global_set:
                        global_set.add((new_a, new_b))
                    incident[w].add((new_a, new_b))
            elif edge[1] == u and edge[0] != v:
                if edge in global_set:
                    global_set.remove(edge)
                    x = edge[0]
                    new_a = min(w, x)
                    new_b = max(w, x)
                    if (new_a, new_b) not in global_set:
                        global_set.add((new_a, new_b))
                    incident[w].add((new_a, new_b))
        # Process incident edges of v
        for edge in list(incident[v]):
            if edge[0] == v and edge[1] != u:
                if edge in global_set:
                    global_set.remove(edge)
                    x = edge[1]
                    new_a = min(w, x)
                    new_b = max(w, x)
                    if (new_a, new_b) not in global_set:
                        global_set.add((new_a, new_b))
                    incident[w].add((new_a, new_b))
            elif edge[1] == v and edge[0] != u:
                if edge in global_set:
                    global_set.remove(edge)
                    x = edge[0]
                    new_a = min(w, x)
                    new_b = max(w, x)
                    if (new_a, new_b) not in global_set:
                        global_set.add((new_a, new_b))
                    incident[w].add((new_a, new_b))
        # Union v into w
        union(w, v)
        # Clear incident lists of u and v
        incident[u].clear()
        incident[v].clear()
        print(len(global_set))

if __name__ == "__main__":
    main()
