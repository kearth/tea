<?php

/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018, Kearth
 */

//默认开启严格模式
declare(strict_types=1);

error_reporting(E_ALL);
//基础常量
define('ROOT_PATH', dirname(__DIR__));

//自动加载
require(ROOT_PATH . "/vendor/autoload.php");
require(ROOT_PATH . "/config/init.php");

//约定大于配置
//引导程序
Tea\Core\Application::bootstrap();
