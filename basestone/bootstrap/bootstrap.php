<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Request;

class Bootstrap
{
    private static $instance = null;

    private function __construct()
    {
    
    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }
    public function run()
    {
        Config::getInstance()->load(CONFIG_PATH);

        //日志系统
        
        //错误处理系统
        
        //路由
        Router::create()->route();

        //分发工作
        Dispatcher::create()->dispatch();
        
        //清理结果
        exit;
    }
}

