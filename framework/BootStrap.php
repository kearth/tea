<?php 

namespace Tea\Framework;

/**
 * 引导类
 */
class BootStrap {

    /**
     * 引导入口
     */
    public static function run() {
        static::init();
        static::process();
    }

    /**
     * 初始化
     */
    private static function init() {
        // 配置初始化
        Config::init(); 
        // 日志初始化
        Logger::init();
        // 错误初始化
        Error::init();
        // 请求初始化 
        Request::init();
        // 路由初始化
        Router::init(); 
        Router::getRouter(); 
    }

    /**
     * 处理流程
     */
    private static function process() {
        // 分发前
        Hook::before();
        // 分发中
        Dispatcher::dispatch();
        // 分发后
        Hook::after();
    }

}
