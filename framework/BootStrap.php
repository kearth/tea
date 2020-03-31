<?php 

namespace Tea\Framework;

/**
 *
 */
class BootStrap {

    /**
     * 引导程序入口
     */
    public static function run() {
        static::init();
        static::process();
    }

    /**
     * 初始化
     */
    private static function init() {
        Config::init(); 
        Logger::init();
        Error::init();
        Router::init(); 
    }

    /**
     * 处理流程
     */
    private static function process() {
        Hook::before();
        Dispatcher::dispatch();
        Hook::after();
    }

}
