<?php
use BaseStone\Bootstrap\Autoload;
use BaseStone\Bootstrap\Bootstrap;

define('ROOT_PATH', realpath(__DIR__."/../"));
define('CONFIG_PATH',ROOT_PATH . '/config/default.ini');


include_once(ROOT_PATH."/basestone/bootstrap/autoload.php");

Autoload::getInstance()->register();
Bootstrap::getInstance()->run();

