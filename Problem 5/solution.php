<?php

$n = 1;

while(1){
    $is_divisible_by_all = true;
    echo $n . PHP_EOL;
    for ($i=1;$i<=20;$i++){
        if($n % $i != 0){
            $is_divisible_by_all = false;
            break;
        }
    }
    if ($is_divisible_by_all){
        break;
    }
    $n++;
}

echo "Found: " . $n . PHP_EOL;