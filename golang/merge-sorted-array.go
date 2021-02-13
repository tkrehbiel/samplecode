// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int)  {
    // Iterate nums2, array of numbers to insert.
    for j := 0; j < n; j++ {
        // Find location where to insert into nums1.
        for i := 0; i <= m; i++ {
            // Insert at where nums1[i] > nums2[j], or end.
            if i == m || nums1[i] > nums2[j] {
                // Shift array to the right.
                copy(nums1[i+1:], nums1[i:])
                // Insert new number.
                nums1[i] = nums2[j]
                // Expand length of nums1 for next iteration.
                m++
                break
            }
        }
    }
}