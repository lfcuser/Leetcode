/*
144 Binary Tree Preorder Traversal
Given the root of a binary tree, return the preorder traversal of its nodes' values.
Constraints:
    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100
*/
class TreeNode {
    constructor(val, left, right) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var preorderTraversal = function(root) {
    if (root === null) {
        return [];
    }
    let nums = [root.val];
    nums = [...nums, ...preorderTraversal(root.left)];
    nums = [...nums, ...preorderTraversal(root.right)];
    return nums;
};

let root = new TreeNode(1);
root.right = new TreeNode(2);
root.right.left = new TreeNode(3);

console.log(preorderTraversal(root), [1,2,3]);

// node ./js/BinaryTreePreorderTraversal.js
