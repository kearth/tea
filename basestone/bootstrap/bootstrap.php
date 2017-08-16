<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Request;
use BaseStone\Core\Singleton;
use BaseStone\Core\Log;
use BaseStone\Core\Container;

class Bootstrap extends Singleton
{
    /**
     *  @name run 运行
     *  @access public
     *  @return void
     */
    public function run()
    {
        //配置初始化
        Config::getInstance()->load(CONFIG_PATH);
        Config::getInstance()->appload();
        
        //路由启动
        Router::getInstance()->run();
        
        //分发启动
        Dispatcher::getInstance()->run();
        
        //结束流程
        Over::getInstance()->run();
    }

}

