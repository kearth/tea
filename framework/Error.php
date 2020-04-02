<?php

namespace Tea\Framework;

/**
 * 错误类
 */
class Error {

    /**
     * 错误信息映射
     */
    public static $errMsgMap = array();

    /**
     * 初始化
     */
    public static function init() : void {
        static::$errMsgMap = Config::initConf(__CLASS__);
    }

    /**
     * 抛异常
     */
    public static function throw(string $message, $code = -1) {
        if (isset(static::$errMsgMap[$code])) {
            $message = static::$errMsgMap[$code];
        }
        throw new \Exception($message, $code); 
    }

    /**
     * 安全调用
     * TODO
     */
    public static function safe() {
        try {
        
        } catch (\Throwable $throwable) {
        
        }
    }

}
