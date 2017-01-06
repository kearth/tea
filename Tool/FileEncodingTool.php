<?php
ini_set('max_execution_time',0);
ini_set('memory_limit','256M');
$path = '/development/git/ksmCMS/Storage/glyphicons-halflings-regular.ttf';
$file = file_get_contents($path);
$filetype = mb_detect_encoding($file,array('utf-8','gbk','latin1','big5'));
if($filetype === 'ISO-8859-1'){
    $file = utf8_encode($file);
}
file_put_contents($path,$file);
