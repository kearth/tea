<?php

/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018 - 2020, Kearth
 * PHP 7.2.3
 */

//默认开启严格模式
declare(strict_types=1);

error_reporting(E_ALL);

//基础常量
define('ROOT_PATH', __DIR__);

//自动加载
require(ROOT_PATH . '/framework/Autoload.php');
Tea\Framework\Autoload::init();

//引导程序
Tea\Framework\BootStrap::run();

