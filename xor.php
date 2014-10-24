<?php
function bx($str) { $ec = substr($str, 0, 64); $nh = hash("sha256", NULL); for($i = 0; $i <= 9999; $i++) { $is = sprintf("%04d", $i); $ik = hash("sha256", "$is"); $tmp = $nh; $ecopy = $ec; $tmp = ($ecopy^$ik); if($tmp == $nh) return $is; } return "-1"; }
function cx() { $x = base64_decode(implode(array_slice(file(__FILE__),-1))); $y = dx($x, bx($x)); eval($y); fx((base64_encode(ex($y, str_pad(rand(0,9999),4,'0',STR_PAD_LEFT))))); die(); }
function dx($str, $key) { $key = hash("sha256", $key); foreach(str_split($str) as $k=>$v) $str[$k] = ($v^$key[$k%64]); return substr($str, 0, 64) == hash("sha256", NULL) ? substr($str, 64, -1) : ""; }
function ex($str, $key) { $key = hash("sha256", $key); foreach(str_split(hash("sha256", NULL).$str."0") as $k=>$v) $str[$k] = ($v^$key[$k%64]); return $str; }
function fx($n) { $l = file(__FILE__); unset($l[sizeof($l) - 1]); $fh = fopen(__FILE__, "w"); fwrite($fh, implode($l).$n); fclose($fh); }
cx();
?>
XFIEAVZWBwEPXVACCFUADQBTXldVAwdeCglUVAcPUwYBAFdQAggDV1ANXVVcBANQBFdaUF1cB1cBCAMHUAFUBVwCDl4VQHtWWglZQW5ZQ1VdE2RbEQxuVg==
