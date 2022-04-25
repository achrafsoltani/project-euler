<?php

declare(strict_types=1);

echo sum_of_multiples(1000, 3,5);

function sum_of_multiples(int $max, int ...$args) : int
{
    $sum = 0;
    for ($i=1;$i<$max;$i++)
    {
        $is_multiplier = false;
        $value = 0;
        foreach ($args as $arg) {
            if ($i % $arg == 0){
                $is_multiplier = true;
                $value = $i;
                break 1;
            }
        }
        $sum += $value;
    }
    return $sum;
}