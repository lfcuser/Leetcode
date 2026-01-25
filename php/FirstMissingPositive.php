<?php
/*
    41. First Missing Positive
    Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
    
    You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.
 */
class Solution
{

    /**
     * @param int[] $nums
     * @return int
     */
    function firstMissingPositive($nums)
    {
        $map = array_flip($nums);
        $res = 1;
        while (isset($map[$res])) {
            $res++;
        }
        return $res;
    }
}

$s = new Solution();

var_dump($s->firstMissingPositive([1,2,0]));
var_dump($s->firstMissingPositive([3,4,-1,1]));
var_dump($s->firstMissingPositive([7,8,9,11,12]));

// php ./php/FirstMissingPositive.php
