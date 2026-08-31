/*
141 Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Constraints:
    The number of the nodes in the list is in the range [0, 104].
    -105 <= Node.val <= 105
    pos is -1 or a valid index in the linked-list.
*/

class ListNode {
    constructor(val, next) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    try {
        JSON.stringify(head);
    } catch (e) {
        return true;
    }
    return false;
};

let first = new ListNode(1);
let second = new ListNode(2);
first.next = second;
second.next = first;
console.log(hasCycle(first));


first = new ListNode(1);
second = new ListNode(2);
first.next = second;
console.log(hasCycle(first));

// node ./js/LinkedListCycle.js
