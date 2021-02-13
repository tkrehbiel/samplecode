// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {

    // Uses a simple array of runes (chars) for the stack.
    stack := make([]rune, len(s))
    stackIndex := 0
    
    // Essentially a stack where each pop must match last push.
    
    // Iterate the input string.
    for _, runeVal := range s {
        switch runeVal {
        case '(', '[', '{':
            // Push each open paren.
			stack[stackIndex] = runeVal
            stackIndex++
        case ')':
            // Compare close paren to top of stack.
            if stackIndex < 1 || stack[stackIndex-1] != '(' {
                   return false
            }
            stackIndex--
        case ']':
            // Compare close bracket to top of stack.
            if stackIndex < 1 || stack[stackIndex-1] != '[' {
                return false
            }
            stackIndex--
        case '}':
            // Compare close brace to top of stack.
            if stackIndex < 1 || stack[stackIndex-1] != '{' {
                return false
            }
            stackIndex--
        default:
            // Any other character is a failure condition.
            return false
        }
    }
    
    // Any leftover unclosed parens is a failure condition.
    if stackIndex > 0 {
        return false
    }
    
    // If we get here all parens must match.
    return true
}
