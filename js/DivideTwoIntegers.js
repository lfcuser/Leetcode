/*
    29 Divide Two Integers

    Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

    The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, 
    and -2.7335 would be truncated to -2.

    Return the quotient after dividing dividend by divisor.

    Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: 
    [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, 
    and if the quotient is strictly less than -231, then return -231.

    Constraints:
        -231 <= dividend, divisor <= 231 - 1
        divisor != 0
*/

/**
 * @param {number} dividend
 * @param {number} divisor
 * @return {number}
 */
var divide = function(dividend, divisor) {
    const max = 2147483647;
    const min = -2147483648;
    if (dividend === min && divisor === -1) {
        return max;
    }
    if (dividend === min && divisor === 1) {
        return min;
    }

    const isNegative = divisor < 0 === dividend < 0;
    dividend = dividend < 0 ? dividend : -dividend;
    divisor = divisor < 0 ? divisor : -divisor;

    let result = 0;
    while (dividend <= divisor) {
        let tempDivisor = divisor;
        let multiple = 1;

        while (dividend <= (tempDivisor << 1) && (tempDivisor >= (min >> 1))) {
            tempDivisor <<= 1;
            multiple <<= 1;
        }

        dividend -= tempDivisor;
        result += multiple;
    }
    return isNegative ? result : -result;
};

console.log(divide(10, 3), 3);
console.log(divide(7, -3), -2);
console.log(divide(-2147483648, -1), 2147483647);
console.log(divide(2147483647, 1), 2147483647);
console.log(divide(-2147483648, 1), -2147483648);
console.log(divide(2147483647, 2), 1073741823);

// node ./js/DivideTwoIntegers.js
