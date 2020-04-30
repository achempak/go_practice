package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i:= range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
             pc[byte(x>>(1*8))] +
             pc[byte(x>>(2*8))] +
             pc[byte(x>>(3*8))] +
             pc[byte(x>>(4*8))] +
             pc[byte(x>>(5*8))] +
             pc[byte(x>>(6*8))] +
             pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var result byte
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func PopCountSlow(x uint64) int {
	result := uint64(0)
	for i := 0; i < 64; i++ {
		result += (x & 1)
		x = x >> 1
	}
	return int(result)
}

func PopCountClearRight(x uint64) int {
	result := 0
	for x > 0 {
		result++
		x &= (x-1)
	}
	return result
}		
