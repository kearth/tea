<?php
use BaseStone\Bootstrap\Autoload;

define('ROOT_PATH',realpath(__DIR__."/../"));

include_once(ROOT_PATH."/basestone/bootstrap/autoload.php");

$autoload = new Autoload();
$autoload->register();

use BaseStone\Bootstrap\Bootstrap;
Bootstrap::getInstance()->run();

