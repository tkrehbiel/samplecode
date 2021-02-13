// https://leetcode.com/problems/valid-anagram/

public class Solution {

    // Build a histogram of letters in a string.
    private Dictionary<char, int> BuildHistogram( string s ) {
        // Uses a Dictionary to store count of letters.
        var h = new Dictionary<char, int>();
        foreach( char c in s ) {
            if( h.ContainsKey(c) )
                h[c] = h[c] + 1;
            else
                h[c] = 1;
        }
        return h;
    }
    
    // Is the second string an anagram of the first?
    public bool IsAnagram(string s, string t) {
        // Edge case: Both inputs null.
        if( s == null || t == null )
            return false;
        
        // If the length of the strings differ, they can't be anagrams.
        if( s.Length != t.Length )
            return false;

        // Build histograms of both strings.
        Dictionary<char, int> histogram1 = BuildHistogram(s);
        Dictionary<char, int> histogram2 = BuildHistogram(t);
        
        // If the histograms are the same, they are anagrams.
        foreach( char key in histogram1.Keys ) {
            // If second string does not contain this letter, 
            // it's not an anagram.
            if( !histogram2.ContainsKey( key ) )
                return false;
            // If second string does not have the same count of letter,
            // it's not an anagram.
            if( histogram2[key] != histogram1[key] )
                return false;
        }
        
        // If we pass all tests above, it's an anagram.
        return true;
    }
}
