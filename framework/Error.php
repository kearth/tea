<?php

namespace Tea\Framework;

/**
 * 错误类
 */
class Error {

    public static $errMsgMap = array();

    public static function init() {
        static::$errMsgMap = Config::get('errMsgMap', 'error');
    }

    public static function throw(string $message, $code = -1) {
        if (isset(static::$errMsgMap[$code])) {
            $message = static::$errMsgMap[$code];
        }
        throw new \Exception($message, $code); 
    }

    public static function safe() {
        try {
        
        } catch (\Throwable $throwable) {
        
        }
    }


}
