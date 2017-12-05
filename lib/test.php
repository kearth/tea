<?php

define('ROOT_PATH', realpath(__DIR__."/../"));
require ROOT_PATH . "/kernel/Autoload.php";

Akf\Kernel\Autoload::register(ROOT_PATH);

new Akf\Kernel\Config();


