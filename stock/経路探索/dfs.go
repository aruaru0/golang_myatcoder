// Node :
type Node struct {
	to []int
}

var route []int

//
// 目的地までのルートを探索するDFS
// 循環（閉路）が無い場合のみ利用可能
// ※何度も使う場合は、routeを初期化すること
//
func dfs(from, to, prev int) bool {
	if to == from {
		route = append(route, to)
		return true
	}
	ret := false
	for _, v := range node[from].to {
		if v == prev {
			continue
		}
		ret = dfs(v, to, from)
		if ret == true {
			break
		}
	}
	if ret == true {
		route = append(route, from)
	}
	return ret
}