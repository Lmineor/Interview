package other

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	last := 1
	for i := 0; i < x; i++ {
		ii := i * i
		if ii == x {
			return i
		} else if ii < x {
			last = i
		} else {
			return last
		}
	}
	return last
}
