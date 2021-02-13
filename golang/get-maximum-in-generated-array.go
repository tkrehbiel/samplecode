// https://leetcode.com/problems/get-maximum-in-generated-array/

// Get maximum value in a generated array of length n.
func getMaximumGenerated(n int) int {
    
    // Edge cases
    switch(n) {
    case 0: 
        return 0
    case 1: 
        return 1
    case 2: 
        return 1
    }
    
    var maxNum int
    
    // Create the array we're going to populate.
    // Make it one bigger to handle overflow.
    nums := make([]int, n+1)
    
    // First two elements always 0 and 1.
    nums[0] = 0
    nums[1] = 1

    // Remainder of elements populated by an odd formula.
    // Loop iterates by 2 so it's half of n rounded up.
    for i := 1; i < (n+1)/2; i++ {
        // Each iteration populates two elements x and y in the array.
        x := i*2
        y := i*2+1
        nums[x] = nums[i]
        nums[y] = nums[i] + nums[i+1]
        // Test for maximums as we go.
        if nums[x] > maxNum {
            maxNum = nums[x]
        }
        if nums[y] > maxNum {
            maxNum = nums[y]
        }
    }
    
    return maxNum
}