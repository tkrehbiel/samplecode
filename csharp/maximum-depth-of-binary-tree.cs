// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution {
    
    // Find maximum depth of a binary tree using recursion.
    private int RecursiveMaxDepth(TreeNode node, int level) {
        if( node == null ) return level;
        // Count current level.
        level++;
        // Explore left and right branches recursively.
        int leftDepth = RecursiveMaxDepth(node.left, level);
        int rightDepth = RecursiveMaxDepth(node.right, level);
        // Count whichever branch was deeper.
        if( leftDepth > level ) level = leftDepth;
        if( rightDepth > level ) level = rightDepth;
        return level;
    }
    
    public int MaxDepth(TreeNode root) {
        return RecursiveMaxDepth(root, 0);
    }
}
