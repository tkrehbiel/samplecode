// https://leetcode.com/problems/simplify-path/

public class Solution {

    // Transform relative path to absolute path.
    public string SimplifyPath(string path) {
        // Uses a stack to keep track of directory depth.
        var stack = new Stack<string>();

        // Split input string on slash.
        string[] parts = path.Split('/');

        // Examine each part of the path.
        for( int i = 0; i < parts.Length; i++ ) {
            string part = parts[i];

            // We may ignore empty strings or . as no-ops.
            if( part == "" || part == "." ) continue;

            // The .. part pops the top back off the stack.
            if( part == ".." ) {
                if( stack.Count > 0 )
                    stack.Pop();
            }
            else {
                // Otherwise the path is pushed onto the stack.
                stack.Push(part);
            }
        }

        // If the stack is empty, it's the root level.
        if( stack.Count == 0 )
            return "/";

        // Would have liked to enumerate the stack object,
        // but had to do so in reverse order which isn't exposed.
        // So we convert it into a simple array of strings.
        string[] stackParts = stack.ToArray();
        
        // Concatenate the stack strings in reverse order,
        // separated by / character.
        var sb = new StringBuilder();
        for( int i = stackParts.Length - 1; i >= 0; i-- ) {
            sb.Append("/");
            sb.Append(stackParts[i]);
        }
        
        return sb.ToString();
    }
}
