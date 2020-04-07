var C [51][51]int // C[n][k] -> nCk

func combTable(N int) {
	C[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}
}