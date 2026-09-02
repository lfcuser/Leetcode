/*
206 Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:

    The number of nodes in the list is the range [0, 5000].
    -5000 <= Node.val <= 5000
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
var reverseList = function(head) {
    if (head === null || head.next === null) {
        return head;
    }
    let current = null;
    while (true) {
        let next = head.next;
        head.next = current;
        current = head;
        if (next === null) {
            break;
        }
        head = next;
    }
    return head;
};

first = new ListNode(1);
second = new ListNode(2);
first.next = second;
console.log(first);
console.log(reverseList(first));

// node ./js/ReverseLinkedList.js