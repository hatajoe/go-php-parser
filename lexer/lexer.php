<?php

function get_token($name, $literal, $line, $col) {
    return [
        "tok" => $name,
        "lit" => $literal,
        "line" => $line,
        "col" => $col,
    ];
}

function update_col($literal, &$col) {
    if (preg_match("/\\n/", $literal)) {
        $col = 1;
    } else {
        $col += strlen($literal);
    }
}

$code = file_get_contents('php://stdin');
$tokens = token_get_all($code);

$line = 1;
$col = 1;
$unset_indices = [];
foreach ($tokens as $idx => $token) {
    if (is_array($token)) {
        $tok_name = token_name($token[0]);
        if ($tok_name === "T_DOUBLE_COLON") {
            $tok_name = "T_PAAMAYIM_NEKUDOTAYIM";
        }
        if (in_array($tok_name, ["T_OPEN_TAG", "T_WHITESPACE"])) {
            $unset_indices[] = $idx;
        }
        $tokens[$idx] = get_token($tok_name, $token[1], $token[2], $col);
        $line = $token[2];
        update_col($token[1], $col);
    } else {
        $tokens[$idx] = get_token($token, $token, $line, $col);
        update_col($token, $col);
    }
}
foreach ($unset_indices as $idx) {
    unset($tokens[$idx]);
}
$tokens = array_values($tokens);
echo json_encode($tokens, JSON_UNESCAPED_SLASHES|JSON_PRETTY_PRINT);
