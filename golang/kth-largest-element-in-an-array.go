// https://leetcode.com/problems/kth-largest-element-in-an-array/

import "sort"

func findKthLargest(nums []int, k int) int {
    // Assumes we can modify the input array.
    // Otherwise a copy will need to be made.
    // Sort the array from lowest to highest.
    sort.Ints(nums)
    // Kth largest element will be at length-k.
    return nums[len(nums)-k]
}