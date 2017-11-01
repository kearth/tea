<?php
/**
 *  入口文件
 */
//开启严格模式
declare(strict_types=1);

use Akf\Core\Kernel\{
    Autoload,
    Config
};


define('ROOT_PATH', realpath(__DIR__."/../"));
define('CONFIG_PATH_ROOT', ROOT_PATH . '/config/');
define('CONFIG_DEFAULT', CONFIG_PATH_ROOT . 'default.php');
define('APPLICATION', ROOT_PATH . '/application/');
define('CONTROLLER', APPLICATION . 'controller/');
define('MODEL', APPLICATION . 'model/');
define('VIEW', APPLICATION . 'view/');
define('ENV_PRODUCT', 'product');
define('ENV_TEST', 'test');
define('ENV_DEV', 'dev');

include_once(ROOT_PATH . "/core/kernel/config.php");
include_once(ROOT_PATH . "/core/kernel/autoload.php");


//加载配置文件
Config::loadDefaultCfg(CONFIG_DEFAULT);

//自动加载注册
Autoload::register(Config::get('alias'));

//容器绑定初始化
Container::init(Config::get('bind'));
Container::initSingleton(Config::get('bindSingle'));

if (ENV === ENV_PRODUCT) {
    error_reporting(0);
} elseif (ENV === ENV_TEST) {
    error_reporting(E_ALL);
} else {
    error_reporting(E_ALL ^ E_NOTICE);
}

//引导程序
Bootstrap::run();
