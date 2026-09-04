/*
404 Sum of Left Leaves

Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Constraints:

    The number of nodes in the tree is in the range [1, 1000].
    -1000 <= Node.val <= 1000
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
 * @return {number}
 */
var sumOfLeftLeaves = function(root) {
    return getLeft(root, false);
};


function getLeft(node, isLeft) {
    let sum = 0;
    if (node === null) {
        return sum;
    }
    if (isLeft && node.right === null && node.left === null) {
        sum += node.val;
    }
    sum += getLeft(node.left, true);
    sum += getLeft(node.right, false);
    return sum;
}


let root = new TreeNode(3);
root.left = new TreeNode(9);
root.right = new TreeNode(20);
root.right.left = new TreeNode(15);

console.log(sumOfLeftLeaves(root), 24);

// node ./js/SumOfLeftLeaves.js
