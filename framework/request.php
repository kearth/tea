<?php

namespace Tea\Framework;

/**
 * 请求类
 */
class Request {

    public static $server = array();
    public static $request = array();

    public static function init() {
        static::$request = $_REQUEST; 
        static::$server = $_SERVER;
        unset($_REQUEST);
        unset($_SERVER);
    }

    public static function getParams() {
        return static::$request;
    }

    public static function getServer() {
        return static::$server;
    }

    public static function getParam($name) {
        if (isset(static::$request[$name])){
            return static::$request[$name];
        } 
        Error::throw('此参数不存在');
    } 

    public static function setParam($name, $value) {
        static::$request[$name] = $value;
    }

}
