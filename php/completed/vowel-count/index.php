<?php
    echo getCount('a');
    function getCount($str) {
        return preg_match_all('/[aeiou]/i',$str,$matches);
    }
?>

