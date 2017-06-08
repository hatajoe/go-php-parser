<?php

function get_token($name, $literal, $line) {
    return [
        "tok" => $name,
        "lit" => $literal,
        "pos" => $line,
    ];
}

$code = file_get_contents('php://stdin');
$tokens = token_get_all($code);

$line = 0;
$unset_indices = [];
foreach ($tokens as $idx => $token) {
    if (is_array($token)) {
        $tok_name = token_name($token[0]);
        if ($tok_name === "T_DOUBLE_COLON") {
            $tok_name = "T_PAAMAYIM_NEKUDOTAYIM";
        }
        if (in_array($tok_name, ["T_OPEN_TAG", "T_WHITESPACE"])) {
            $unset_indices[] = $idx;
            continue;
        }
        $tokens[$idx] = get_token($tok_name, $token[1], $token[2]);
        $line = $token[2];
    } else {
        $tokens[$idx] = get_token($token, $token, $line);
    }
}
foreach ($unset_indices as $idx) {
    unset($tokens[$idx]);
}
$tokens = array_values($tokens);
echo json_encode($tokens, JSON_UNESCAPED_SLASHES|JSON_PRETTY_PRINT);
