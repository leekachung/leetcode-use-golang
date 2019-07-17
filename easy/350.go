func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums := make([]int, 0)
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); i++ {
		if nums1[i] > nums2[j] {
			for {
				j++
				if j == len(nums2) || nums1[i] <= nums2[j] {
					break
				}
			}
			if j == len(nums2) {
				break
			}
		}
		if nums1[i] == nums2[j] {
			nums = append(nums, nums1[i])
			j++
		}
	}
	return nums
}
