/*
3452 Sum of Good Numbers

Given an array of integers nums and an integer k, an element nums[i] is considered good if it is strictly greater than the elements at indices i - k and i + k (if those indices exist). If neither of these indices exists, nums[i] is still considered good.

Return the sum of all the good elements in the array.

Constraints:

    2 <= nums.length <= 100
    1 <= nums[i] <= 1000
    1 <= k <= floor(nums.length / 2)
*/
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var sumOfGoodNumbers = function(nums, k) {
    let sum = 0;
    for (let i = 0; i < nums.length; i++) {
        if ((nums[i-k] === undefined || nums[i-k] < nums[i])
            && (nums[i+k] === undefined || nums[i+k] < nums[i])
        ) {
            sum += nums[i];
        }
    }
    return sum;
};

console.log(sumOfGoodNumbers([1,3,2,1,5,4], 2), 12);
console.log(sumOfGoodNumbers([2,1], 1), 2);

// node ./js/SumOfGoodNumbers.js
