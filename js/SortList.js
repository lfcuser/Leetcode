/*
148 Sort List

Given the head of a linked list, return the list after sorting it in ascending order.

Input: head = [4,2,1,3]
Output: [1,2,3,4]

Constraints:

    The number of nodes in the list is in the range [0, 5 * 104].
    -105 <= Node.val <= 105
*/

class ListNode {
    constructor(val, next) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var sortList = function(head) {
    let arr = listToArray(head);
    if (arr.length === 0) {
        return null;
    }
    if (arr.length === 1) {
        return new ListNode(arr[0]);
    }
    arr = mergeSort(arr);
    return fillList(arr, 0, new ListNode());
};

function mergeSort(arr) {
    if (arr.length === 1) {
        return arr;
    }

    const mid = Math.trunc(arr.length / 2);
    let left = arr.slice(0, mid);
    let right = arr.slice(mid);

    left = mergeSort(left);
    right = mergeSort(right);

    let res = [];
    let i = 0;
    let j = 0;

    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            res.push(left[i++]);
        } else {
            res.push(right[j++]);
        }
    }
    while (i < left.length) {
        res.push(left[i++]);
    }
    while (j < right.length) {
        res.push(right[j++]);
    }

    return res;
}

function fillList(arr, i, head) {
    if (i >= arr.length) {
        return null;
    }

    head.val = arr[i];
    head.next = fillList(arr, i + 1, new ListNode());
    return head;
}

function listToArray(head) {
    let arr = [];

    while (head !== null) {
        arr.push(head.val);
        head = head.next;
    }

    return arr;
}

head = fillList([4,2,1,3], 0, new ListNode());
console.log(listToArray(head));
console.log(listToArray(sortList(head)));

// node ./js/SortList.js
