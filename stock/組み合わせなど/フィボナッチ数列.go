const mod = int(1e9 + 7)

func fib(n int) int {
	a, b, c, d := 1, 1, 1, 0
	e, f, g, h := 1, 0, 1, 0
	n = n - 1
	for n > 0 {
		if n%2 == 1 {
			i := a*e%mod + b*g%mod
			j := a*f%mod + b*f%mod
			k := c*e%mod + d*g%mod
			l := c*f%mod + d*h%mod
			e, f, g, h = i%mod, j%mod, k%mod, l%mod
		}
		i := a*a%mod + b*c%mod
		j := a*b%mod + b*d%mod
		k := c*a%mod + d*c%mod
		l := c*b%mod + d*d%mod
		a, b, c, d = i%mod, j%mod, k%mod, l%mod
		n /= 2
	}
	return e + f
}