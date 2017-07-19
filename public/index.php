<?php
use Bootstrap\Bootstrap;

define('ROOT',realpath(__DIR__."/../"));

include_once(ROOT."/bootstrap/bootstrap.php");

Bootstrap::getInstance()->run();

