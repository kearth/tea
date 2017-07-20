<?php
use Bootstrap\Bootstrap;

define('ROOT',realpath(__DIR__."/../"));

include_once(ROOT."/bootstrap/bootstrap.php");

include_once(ROOT."/bootstrap/autoload.php");
$bootstrap = Bootstrap::getInstance();
$bootstrap->run();


