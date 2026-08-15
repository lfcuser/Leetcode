// 704 Binary Search

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let left = 0;
    let right = nums.length;

    while (left <= right) {
        current = left + Math.trunc((right - left) / 2);

        if (nums[current] === target) {
            return current;
        }
        if (nums[current] < target) {
            left = current + 1;
        } else {
            right = current - 1;
        }
    }

    return -1;
};

console.log(search([-1, 0, 3, 5, 9, 12], 9));

// node ./js/BinarySearch.js