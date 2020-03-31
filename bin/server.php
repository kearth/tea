<?php 
$dir = __DIR__;
$cmd = "php -S localhost:8000  $dir/../index.php";
shell_exec($cmd);
