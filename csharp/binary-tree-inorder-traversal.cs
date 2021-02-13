// https://leetcode.com/problems/binary-tree-inorder-traversal/
// (Had to look up what "inorder traversal" meant.)

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
    private List<int> nums = new List<int>();
    
    private void RecursiveTraversal(TreeNode node) {
        if( node == null ) return;
        // "Inorder traversal" means left child, then current, then right child.
        // (As opposed to "preorder" and "postorder" traversal.)
        RecursiveTraversal(node.left);
        nums.Add(node.val);
        RecursiveTraversal(node.right);
    }

    public IList<int> InorderTraversal(TreeNode root) {
        RecursiveTraversal(root);
        return nums;
    }
}