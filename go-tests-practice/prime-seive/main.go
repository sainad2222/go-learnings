package prime_seive

func BrutePrimality(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SeivePrimality(n int) bool {
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += 1
	}
	return true
}