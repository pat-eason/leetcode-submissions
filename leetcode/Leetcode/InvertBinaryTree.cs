using System;
namespace leetcode.Leetcode
{
    public class InvertBinaryTree
    {
        public static TreeNode InvertTree(TreeNode root)
        {
            if (root == null)
            {
                return null;
            }

            TreeNode leftNode = InvertTree(root.left);
            TreeNode rightNode = InvertTree(root.right);

            root.left = rightNode;
            root.right = leftNode;

            return root;
        }
    }
}

