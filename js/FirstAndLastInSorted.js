/*
    34. Find First and Last Position of Element in Sorted Array
    
    Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

    If target is not found in the array, return [-1, -1].
 */

var searchRange = function (nums, target) {
    let left = binarySearch(nums, target);
    let right = binarySearch(nums, target, false);
    return [left, right];
};

function binarySearch(arr, target, mostLeft = true) {
    let left = 0;
    let right = arr.length - 1;
    
    let current = 0;
    while (left <= right) {
        current = left + Math.floor(right - left / 2);

        if (arr[current] === target) {
            if (mostLeft && arr[current - 1] === target) {
                right = current - 1;
            } else if (!mostLeft && arr[current + 1] === target) {
                left = current + 1;
            } else {
                return current;
            }
        }
        if (arr[current] < target) {
            left = current + 1;
        } else {
            right = current - 1;
        }
    }
    
    return -1;
}

console.log(searchRange([5, 7, 7, 8, 8, 10], 8));
console.log(searchRange([5, 7, 7, 8, 8, 10], 6));
console.log(searchRange([], 0));
console.log(searchRange([5, 7, 7, 8, 8, 10], 6));
console.log(searchRange([1], 1));
console.log(searchRange([3,3,3], 3));
