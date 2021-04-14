func strWithout3a3b(A int, B int) string {
	var ans string
	a, b := "a", "b"

	if B > A {
		A, B = B, A
		a, b = b, a
	}

	for A > 0 || B > 0 {
		if A > 0 {
			ans += a
			A--
		}

		if A > B {
			ans += a
			A--
		}

		if B > 0 {
			ans += b
			B--
		}
	}

	return ans
}