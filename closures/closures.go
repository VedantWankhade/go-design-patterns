package closures

func addN(m int) func(n int) int {
	return func(n int) int {
		r := m + n
		m++
		return r
	}
}
