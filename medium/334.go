func increasingTriplet(nums []int) bool {
    one := 2 << 31 - 1
    two := one
    for _, v := range nums {
        if v <= one {
            one = v
        } else if v <= two {
            two = v
        } else {
            return true
        }
    }
    return false
}
