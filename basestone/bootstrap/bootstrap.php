<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Request;
use BaseStone\Core\Log;

class Bootstrap
{
    private static $instance = null;

    private function __construct(){}

    private function __clone(){}

    private function __wakeup(){}

    private function __sleep(){}

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    /**
     *  @name run 运行
     *  @access public
     *  @return void
     */
    public function run()
    {
        //配置初始化
        Config::getInstance()->init(CONFIG_PATH);

        //日志监控初始化
        Log::getInstance()->init();
        
        //路由启动
        Router::create()->run();
        
        //分发启动
        Dispatcher::create()->run();
        
        //结束流程
        Over::create()->run();
    }
}

