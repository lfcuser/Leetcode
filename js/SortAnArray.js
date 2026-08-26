/*
912 Sort an Array

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Constraints:

    1 <= nums.length <= 5 * 104
    -5 * 104 <= nums[i] <= 5 * 104
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function(nums) {
    if (nums.length === 1) {
        return nums;
    }

    const mid = Math.trunc(nums.length / 2);
    let left = nums.slice(0, mid);
    let right = nums.slice(mid);
    left = sortArray(left);
    right = sortArray(right);

    const res = [];
    let i = 0;
    let j = 0;
    while ((left[i] ?? null) !== null && (right[j] ?? null) !== null) {
        if (left[i] <= right[j]) {
            res.push(left[i++]);
        } else {
            res.push(right[j++]);
        }
    }
    while ((left[i] ?? null) !== null) {
        res.push(left[i++]);
    }
    while ((right[j] ?? null) !== null) {
        res.push(right[j++]);
    }
    return res;
};

console.log(sortArray([5,2,3,1]));
console.log(sortArray([5,1,1,2,0,0]));

// node ./js/SortAnArray.js
