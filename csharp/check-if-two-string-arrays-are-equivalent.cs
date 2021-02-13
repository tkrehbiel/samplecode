// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

public class Solution {
    
    // Concatenate all parts of the word array.
    private StringBuilder ConcatenateWords( string[] word ) {
        // Using StringBuilder because normal string concatenation is expensive.
        var sb = new StringBuilder();
        foreach( string s in word ) {
            sb.Append( s );
        }
        return sb;
    }
    
    public bool ArrayStringsAreEqual(string[] word1, string[] word2) {
        // Concatenate both arrays with StringBuilders.
        var sb1 = ConcatenateWords( word1 );
        var sb2 = ConcatenateWords( word2 );
        // Now we need only compare the two results.
        return sb1.ToString().Equals( sb2.ToString() );
    }
}