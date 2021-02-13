// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/

// Return concatenated binary of all numbers from 1 to n
func concatenatedBinary(n int) int {

    const constraint = 1000000007
    
    // Initializers
    currentValue := 1       // current index starts at 1
    significantBits := 2    // significant bits required for current value (for 2 and 3)
    thresholdValue := 4     // threshold at which significant bits change
    
    // Begin enumeration at 2
    for i := 2; i <= n; i++ {
        if i >= thresholdValue {
            significantBits++
            thresholdValue *= 2
        }
        // Shift value left to make room for new value.
        currentValue <<= significantBits
        // Add new value at rightmost.
        currentValue += i
        // Must take care to stay within the constraint during processing
        // otherwise we will overflow 32-bit or 64-bit integers by a lot.
        if currentValue > constraint {
           currentValue = currentValue % constraint 
        }
    }
    
    return currentValue % constraint
}
