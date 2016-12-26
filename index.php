<?php
/**
 * 文件名：index.php
 * 功能：
 * 1.定义常量
 * 2.加载类库
 * 3.初始运行
 * @author:jiakun<kunsama@163.com>
 * @version:1.0
 */

/**********************定义常量*************/
define('ROOT',realpath(__DIR__));           //根目录
define('CONTROLLER',ROOT.'/Controller');    //控制器目录
define('CORE',ROOT.'/Core');                //核心目录
define('EXTENSION',ROOT.'/Extension');      //扩展目录
define('LIB',ROOT.'/Lib');                  //公共库目录
define('MODEL',ROOT.'/Model');              //模型目录
define('STORAGE',ROOT.'/Storage');          //资源目录
define('TEST',ROOT.'/Test');                //测试目录
define('VIEW',ROOT.'/View');                //视图目录
define('VENDOR',ROOT.'/vendor');            //Composer依赖目录
define('CONFIG',ROOT.'/conf.php');          //配置文件
define('DEBUG',true);                       //调试模式：true开启，false关闭


/**********************加载类库*************/
require(CORE."/Init.php");
require(ROOT."/vendor/autoload.php");
spl_autoload_register('\core\Init::autoLoad'); //自动加载类

//模式判定
if(DEBUG){
    error_reporting(E_ALL);
    $whoops = new \Whoops\Run();
    $handler = new \Whoops\Handler\PrettyPageHandler();
    $whoops->pushHandler($handler);
    $whoops->register();

} else {
    error_reporting(0);
}


/**********************初始运行*************/
\core\Init::run();
