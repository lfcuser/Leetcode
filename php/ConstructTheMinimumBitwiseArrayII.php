<?php
/*
    3315. Construct the Minimum Bitwise Array II
    
    You are given an array nums consisting of n Prime integers.
    
    You need to construct an array ans of length n, such that, for each index i, the bitwise OR of ans[i] and ans[i] + 1 is equal to nums[i], i.e. ans[i] OR (ans[i] + 1) == nums[i].
    
    Additionally, you must minimize each value of ans[i] in the resulting array.
    
    If it is not possible to find such a value for ans[i] that satisfies the condition, then set ans[i] = -1.
*/
class Solution
{
    /**
     * @param int[] $nums
     * @return int[]
     */
    function minBitwiseArray($nums)
    {
        $remember = [];
        $n = count($nums);
        $ans = array_fill(0, $n, -1);
        
        for ($i = 0; $i < $n; $i++) {
            if ($nums[$i] % 2 === 0) {
                continue;
            }
            
            if (isset($remember[$nums[$i]])) {
                $ans[$i] = $remember[$nums[$i]];
                continue;
            }
            
            $variant = $this->getVariant($nums[$i]);
            if (!is_null($variant) && ($variant | ($variant + 1)) === $nums[$i]) {
                $ans[$i] = $variant;
                $remember[$nums[$i]] = $variant;
            }
        }
        return $ans;
    }
    
    /**
     * When j or j + 1 equals $num, it means the value differs by one bit.
     * The minimum such difference is achieved only for the first "01"
     * bit pattern found from the right.
     */
    private function getVariant(int $num)
    {
        $ibin = decbin($num);
        $ibin = str_pad($ibin, 32, '0', STR_PAD_LEFT);
        $len = strlen($ibin);
        $template = $ibin;
        
        for ($j = $len - 1; $j >= 0; $j--) {
            if ($ibin[$j] === '0' && ($ibin[$j + 1] ?? '0') === '1') {
                $template[$j + 1] = '0';
                return bindec($template);
            }
        }
        
        return null;
    }
}

$s = new Solution();

$inputs = [
    [
        'expect' => [-1,1,4,3],
        'input' => [2,3,5,7],
    ],
    [
        'expect' => [9,12,15],
        'input' => [11,13,31],
    ],
    [
        'expect' => [9,12,15,9,12,15,9,12,15,9,12,15,9,12,15,9,12,15],
        'input' => [11,13,31,11,13,31,11,13,31,11,13,31,11,13,31,11,13,31],
    ],
    [
        'expect' => [-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1],
        'input' =>  [0,2,4,6,8,10,12,14,16,18,20],
    ],
    [
        'expect' => [884532609,741533368,868936608,816315815,150570780,346594696,334726180,921762640,89355880,403165728,931242732],
        'input' => [884532611,741533369,868936609,816315823,150570781,346594697,334726181,921762641,89355881,403165729,931242733],
    ],
    [
        'expect' => [4095,9801,7457,8692,9863,4228,1032,7185,1732,8705,2968,8368,7560,2841,9208,9696],
        'input' => [8191,9803,7459,8693,9871,4229,1033,7187,1733,8707,2969,8369,7561,2843,9209,9697],
    ],
];

foreach ($inputs as $input) {
    // echo "Expect:\n";
    // var_dump($input['expect']);
    // echo "Result:\n";
    $res = $s->minBitwiseArray($input['input']);
    //var_dump($res);
    var_dump($res === $input['expect']);
}

// php ./php/ConstructTheMinimumBitwiseArrayII.php
