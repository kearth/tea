<?php

/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018, Kearth
 * PHP 7.2.3
 */

//默认开启严格模式
declare(strict_types=1);

error_reporting(E_ALL);

//基础常量
define('ROOT_PATH', dirname(__DIR__));
define('AUTOLOAD_PATH', ROOT_PATH . '/vendor/autoload.php');
define('INIT_CONFIG_PATH', ROOT_PATH . '/config/init.php');

//自动加载
require(AUTOLOAD_PATH);

//引导程序
Tea\Core\Application::bootstrap();

