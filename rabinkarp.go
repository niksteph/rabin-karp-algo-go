package rabinkarp

func find(T, P string) int {
	d, q := 256, 19
	n, m := len(T), len(P)
	if m > n {
		return -1
	}
	var p, t int
	for i := 0; i < m; i++ {
		p = (d*p+int(P[i]))%q
		t = (d*t+int(T[i]))%q
	}
	for s := 0; s <= n-m; s++ {
		if t == p && T[s:s+m] == P {
			return s
		}
		if s < n-m {
			t = rollHash(t, T[s], T[s+m], d, q, m)
		}
	}
	return -1
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func expMod(a, b, n int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		d := expMod(a, b/2, n)
		return (d * d) % n
	} else {
		d := expMod(a, b-1, n)
		return (a * d) % n
	}
}

func hash(s string, base, mod int) int {
	if len(s) == 1 {
		return int(s[0]) % mod
	}
	return (int(s[len(s)-1])%mod + base*hash(s[:len(s)-1], base, mod)) % mod
}

func rollHash(prev int, oldHigh, newLow byte, d, n, order int) int {
	h := expMod(d, order-1, n)
	return mod(d*(prev-int(oldHigh)*h)+int(newLow), n)
}
