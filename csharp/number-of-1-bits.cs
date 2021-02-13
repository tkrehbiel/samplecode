// https://leetcode.com/problems/number-of-1-bits/

public class Solution {
    
    // Slower implementation that counts each bit.

    private int HammingWeightNoLookup(uint n) {
        int count = 0;
        while( n != 0 )
        {
            // If rightmost bit is 1, count it.
            if( (n & 1) != 0 ) count++;
            // Bit shift one place to the right.
            n >>= 1;
        }
        return count;
    }

    // Faster implementation using a lookup table.
    // Size of the lookup table could be tuned up or down
    // depending on available memory.
    // eg. Table could be 512 entries and the bit shift would become 9.
    // Or the table could be 128 entries and the bit shift would become 7.

    private bool isLookupInitialized = false;
    private int[] lookup = new int[256];

    private int HammingWeightWithLookup(uint n) {

        // Initialize the lookup table the first time called.
        if( !isLookupInitialized ) {
            for( uint i = 0; i < 256; i++ ) {
                lookup[i] = HammingWeightNoLookup(i);
            }
            isLookupInitialized = true;
        }

        int count = 0;
        while( n != 0 )
        {
            // Lookup value of rightmost 8 significant bits.
            count += lookup[n & 255];
            // Shift number to the right by 8 bits.
            n >>= 8;
        }
        return count;
    }

    // Return the number of 1 bits in a number.
    public int HammingWeight(uint n) {
        return HammingWeightWithLookup(n);
   }
}