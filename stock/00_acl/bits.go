import "math/bits"

func ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}
func bsf(n uint) int {
	return bits.TrailingZeros(n)
}
