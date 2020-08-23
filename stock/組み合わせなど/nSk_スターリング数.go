//
// スターリング数
//   区別できるn個を、区別しないkグループに分ける
// 組み合わせの数。０個のグループは許さない
//
const tblsize = 1000
const mod = 1000000007

var sTable [tblsize][tblsize]int

func initSTable() {
	sTable[0][0] = 0
	sTable[1][0] = 0
	sTable[1][1] = 1
	for n := 2; n < tblsize; n++ {
		sTable[n][0] = 0
		for k := 1; k <= n; k++ {
			sTable[n][k] = sTable[n-1][k-1] + (k*sTable[n-1][k])%mod
			sTable[n][k] %= mod
		}
	}
}

func nSk(n, k int) int {
	if n >= tblsize || k >= tblsize {
		panic("nSk size overflow")
	}
	return sTable[n][k]
}
