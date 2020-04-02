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

//报错等级设置
error_reporting(E_ALL);

/**
 * 基础常量
 */
// 项目目录
define('ROOT_PATH', __DIR__);
// 配置目录
define('CONF_ROOT', ROOT_PATH . '/conf/');
// 默认配置key
define('DEFAULT_CONF_KEY', 'init');
// 成功
define('CODE_SUCCESS', 0);
// 未知错误
define('CODE_UNKNOWN', 999);
// 框架错误
define('CODE_FRAMEWORK', 1);
// 生产环境
define('ENV_ONLINE', 'ONLINE');
// 沙盒环境
define('ENV_SANDBOX', 'SANDBOX');
// 测试环境
define('ENV_TEST', 'TEST');
// 开发环境
define('ENV_DEV', 'DEV');

//自动加载
require(ROOT_PATH . '/framework/Autoload.php');
Tea\Framework\Autoload::init();

//引导程序
Tea\Framework\BootStrap::run();

