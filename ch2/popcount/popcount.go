package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount : count 1bit in x
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

// PopCount2 : count 1bit in x using for loop
func PopCount2(x uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc[byte(x>>uint(i*8))]
	}
	return int(res)
}

// PopCount3 : count 1bit in x using shift
func PopCount3(x uint64) int {
	var res uint64
	for i := 0; i < 64; i++ {
		res += x & 1
		x = x >> 1
	}
	return int(res)
}

// PopCount4 : count 1bit in x using x&(x-1)
func PopCount4(x uint64) int {
	var res int
	for x > 0 {
		res++
		x = x & (x - 1)
	}
	return res
}
