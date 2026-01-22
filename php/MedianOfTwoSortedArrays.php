<?php
/*
    4. Median of Two Sorted Arrays

    Given two sorted arrays nums1 and nums2 of size m and n respectively,
    return the median of the two sorted arrays.
    The overall run time complexity should be O(log (m+n)).
 */
class Solution
{
    /**
     * @param int[] $nums1
     * @param int[] $nums2
     * @return float
     */
    function findMedianSortedArrays(array $nums1, array $nums2): float
    {
        $merged = $this->merge($nums1, $nums2);

        $midIdx = count($merged) / 2;
        $middle = 0;
        if (is_float($midIdx)) {
            $middle = $merged[(int) $midIdx];
        } else {
            $middle = ($merged[$midIdx] + $merged[$midIdx-1]) / 2;
        }
        return $middle;
    }
    
    function merge(array $nums1, array $nums2): array
    {
        $merged = [];
        $i = $j = 0;
        while (isset($nums1[$i]) && isset($nums2[$j])) {
            if ($nums1[$i] <= $nums2[$j]) {
                $merged[] = $nums1[$i++];
            } else {
                $merged[] = $nums2[$j++];
            }
        }

        while (isset($nums1[$i])) {
            $merged[] = $nums1[$i++];
        }

        while (isset($nums2[$j])) {
            $merged[] = $nums2[$j++];
        }
        
        return $merged;
    }
}

$s = new Solution();

$inputs = [
    [
        'expect' => 2.0,
        'input' => ['nums1' => [1,3], 'nums2' => [2]],
    ],
    [
        'expect' => 2.5,
        'input' => ['nums1' => [1,2], 'nums2' => [3,4]],
    ],
    [
        'expect' => 10.0,
        'input' => ['nums1' => [1,2,7,19], 'nums2' => [5,10,14,15,18]],
    ],
    [
        'expect' => 8.5,
        'input' => ['nums1' => [1,2,7,19], 'nums2' => [5,10,14,15]],
    ],
];

foreach ($inputs as $input) {
    // echo "Expect:\n";
    // var_dump($input['expect']);
    // echo "Result:\n";
    $res = $s->findMedianSortedArrays($input['input']['nums1'], $input['input']['nums2']);
    //var_dump($res);
    var_dump($res === $input['expect']);
}

// php ./php/MedianOfTwoSortedArrays.php
