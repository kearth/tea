<?php
/**
 * 入口文件
 * 1.定义常量
 * 2.加载类库
 * 3.初始运行
 **/

define('ROOT',realpath(__DIR__));
define('CONTROLLER',ROOT.'/Controller');
define('CORE',ROOT.'/Core');
define('EXTENSION',ROOT.'/Extension');
define('LIB',ROOT.'/Lib');
define('MODEL',ROOT.'/Model');
define('STORAGE',ROOT.'/Storage');
define('TEST',ROOT.'/Test');
define('VIEW',ROOT.'/View');
define('CONFIG',ROOT.'/Config.php');
define('DEBUG',true);

if(DEBUG){
    error_reporting(E_ALL);
} else {
    error_reporting(0);
}

require(CORE."/Init.php");

spl_autoload_register('\core\Init::autoLoad');

\core\Init::run();
