<?php
/**
 * Tea Framework
 * Version 1.0.0
 *
 * Copyright 2018, Kearth
 */

//默认开启严格模式
declare(strict_types=1);

//入口文件 -> 懒加载配置文件 -> 自动加载文件 -> 引导程序 -> 路由 -> 启动应用处理 -> 处理结束结束请求

use Tea\Kernel\Config;
use Tea\Kernel\Autoload;
use Tea\Kernel\Container;

//基础常量
define('ROOT_PATH', realpath(__DIR__."/../"));
define('INIT_CONFIG_PATH', ROOT_PATH . "/config/init.php");

include_once(ROOT_PATH . "/kernel/Config.php");
include_once(ROOT_PATH . "/kernel/Autoload.php");

//加载配置文件

Config::initLoad(INIT_CONFIG_PATH);

//自动加载注册
Autoload::register();

//容器绑定初始化

new Container();


exit;
Container::init(Config::get('bind'));
Container::initSingleton(Config::get('bindSingle'));



//引导程序
Bootstrap::run();
