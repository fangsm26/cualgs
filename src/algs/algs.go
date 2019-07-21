package algs

// FindMedianSortedArrays finds the median of two sorted integer arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return FindMedianSortedArrays(nums2, nums1)
	}
	min_index, max_index := 0, len1
	for min_index <= max_index {
		i := (min_index + max_index) / 2
		j := (len1+len2+1)/2 - i
		if i < len1 && j > 0 && nums2[j-1] > nums1[i] {
			min_index = i + 1
		} else if i > 0 && j < len2 && nums1[i-1] > nums2[j] {
			max_index = i - 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = Max(nums1[i-1], nums2[j-1])
			}
			if (len1+len2)&1 == 1 {
				return float64(leftMax)
			}
			rightMin := 0
			if i == len1 {
				rightMin = nums2[j]
			} else if j == len2 {
				rightMin = nums1[i]
			} else {
				rightMin = Min(nums1[i], nums2[j])
			}
			return float64(leftMax+rightMin) / 2
		}
	}
	return 0.0
}
