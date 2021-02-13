// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/

import "strings"

var letterLookup [26]string = [26]string{ "a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

func getSmallestString(n int, k int) string {
    
    // Temporary storage in an integer array.
    intArray := make([]int, n)
    
    // Iterate starting at the right.
    for outputIndex := n; outputIndex >= 1; outputIndex-- {
        // Get the highest character value that is 
        // less than or equal to current value of k...
        // ...less the number of digits to the left of the current position
        maxValue := k - (outputIndex-1)
        if maxValue > 26 {
            maxValue = 26
        }
        // Store the value in the array.
        intArray[outputIndex-1] = maxValue
        // Subtract that value from k.
        k -= maxValue
        // Move one position to the left, repeat.
    }
 
    // Use a string builder object because string concatenation is expensive.
    var b strings.Builder
    b.Grow(n)
    
    // Build string in reverse from int array.
    for i := 0; i < n; i++ {
        // Convert value to a character.
        // Uses a simple lookup for this problem.
        c := letterLookup[intArray[i]-1]
        b.WriteString(c)
    }
    
    return b.String()
}