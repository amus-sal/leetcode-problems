var mem map[int]float64

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}
	mem = make(map[int]float64)
	t := int(math.Abs(float64(n)))
	mem[0] = 1
	mem[1] = x
	getPow(x, t)

	if n < 0 {
		return 1 / mem[t]
	}
	return mem[n]
}

func getPow(x float64, n int) {

	if n == 1 {
		mem[n] = x
		return
	}

	if _, out := mem[n]; out {
		return
	}
	if _, ok := mem[n/2]; ok {
		if _, ok2 := mem[n-n/2]; ok2 {
			mem[n] = mem[n-n/2] * mem[n/2]
			return
		}
		getPow(x, n-n/2)
		mem[n] = mem[n-n/2] * mem[n/2]
		return
	}
	getPow(x, n/2)
	getPow(x, n-n/2)
	mem[n] = mem[n-n/2] * mem[n/2]
	return

}