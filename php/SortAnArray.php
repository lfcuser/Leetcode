<?php
/*
912 Sort an Array

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Constraints:

    1 <= nums.length <= 5 * 104
    -5 * 104 <= nums[i] <= 5 * 104
*/

class Solution
{
    /**
     * @param int[] $nums
     *
     * @return int[]
     */
    public function sortArray(array $nums): array
    {
        $length = count($nums);
        if ($length === 1) {
            return $nums;
        }

        $mid = intdiv($length, 2);
        $left = array_slice($nums, 0, $mid);
        $right = array_slice($nums, $mid);

        $left = $this->sortArray($left);
        $right = $this->sortArray($right);

        $res = [];
        $i = $j = 0;
        while (isset($left[$i]) && isset($right[$j])) {
            if ($left[$i] <= $right[$j]) {
                $res[] = $left[$i++];
            } else {
                $res[] = $right[$j++];
            }
        }
        while (isset($left[$i])) {
            $res[] = $left[$i++];
        }
        while (isset($right[$j])) {
            $res[] = $right[$j++];
        }

        return $res;
    }
}

$obj = new Solution();
var_dump($obj->sortArray([5,2,3,1]));
var_dump($obj->sortArray([5,1,1,2,0,0]));

// php ./php/SortAnArray.php
