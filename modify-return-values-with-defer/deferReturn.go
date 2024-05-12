package modifyreturnvalueswithdefer

func addTwo(m int) int {
	return m + 2
}

func addTwoDefer(m int) (res int) {
	res = 99
	defer func() {
		res++
	}()
	return m + 2
}
