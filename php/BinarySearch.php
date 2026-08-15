<?php

// 704 Binary Search

/**
 * @param Integer[] $nums
 * @param Integer $target
 * @return Integer
 */
function search(array $nums, int $target): int
{
    $left = 0;
    $right = count($nums) - 1;

    while ($left <= $right) {
        $current = $left + intdiv($right - $left, 2);
        if ($nums[$current] === $target) {
            return $current;
        }
        if ($nums[$current] < $target) {
            $left = $current + 1;
        } else {
            $right = $current - 1;
        }
    }

    return -1;
}

var_dump(search([-1, 0, 3, 5, 9, 12], 9));

// php ./php/BinarySearch.php