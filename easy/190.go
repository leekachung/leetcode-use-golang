// solution 1 4ms
func reverseBits(num uint32) uint32 {
    var res uint32
    i := 32
    for i > 0 {
        res <<= 1
        res += num & 1
        num >>= 1
        i--
    }
    return res
}

// solution 2 0ms
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; {
		ans <<= 1
		ans |= num&1
		num >>= 1
		i++
	}
	return ans
}
