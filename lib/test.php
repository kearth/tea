<?php
namespace Akf\Kernel;

use Akf\Lib\a;

define('ROOT_PATH', realpath(__DIR__."/../"));
require ROOT_PATH . "/kernel/Autoload.php";

Autoload::register(ROOT_PATH);

print_r(new a());


