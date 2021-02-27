<?php
function findFirstStringInBracket($str) {

    if(strlen($str) > 0) {
        $firstbracket = strstr($str, '(');
        if($firstbracket) {
            $firstbracket = ltrim($firstbracket, '(');
            return strstr($firstbracket, ')', true);
        }
        return '';
    }
    return '';
}

echo findFirstStringInBracket("Muhammad Nur )))(Fadli)");
?>