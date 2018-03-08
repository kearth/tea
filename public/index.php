<?php
/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018, Kearth
 */

//默认开启严格模式
declare(strict_types=1);

//基础常量
define('ROOT_PATH', dirname(__DIR__));

//自动加载
include(ROOT_PATH . "/vendor/autoload.php");

//引导程序
\Tea\Core\Bootstrap::run();
