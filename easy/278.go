// Forward declaration of isBadVersion API.
func isBadVersion(version) bool {}
func firstBadVersion(n int) int {
    left, right := 1, n
    if isBadVersion(left) { return 1 }
    for left < right {
        mid := left + (right - left) / 2
        if !isBadVersion(mid - 1) && isBadVersion(mid) { 
            return mid 
        } else if !isBadVersion(mid) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    if !isBadVersion(left - 1) && isBadVersion(left) { return mid }
    return -1
}
