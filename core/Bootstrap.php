<?php

namespace Tea\Core;
//入口文件 -> 引导程序 -> 配置-> 路由 -> 启动应用处理 -> 处理结束结束请求
class Bootstrap
{
    public static function run() : void
    {
        Config::init(ROOT_PATH . "/config/app.php");

        Env::init(Config::getConfig("env"));

        $component = Config::getConfig("component");
        $map = Config::getConfig("map");
        
        //Container
        //Route
        //Dispatch
        //exit
    }
}
