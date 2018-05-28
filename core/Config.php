<?php

namespace Tea\Core;

class Config
{
    public static $appConfig = [];
    public static $routeConfig = [];

    public static function init()
    {
        static::$appConfig = static::load(APP_CONFIG);
        static::$routeConfig = static::load(ROUTE_CONFIG);
        var_export(static::$appConfig);
    }

    public static function load(string $fileName)
    {
        if (is_file($fileName)) {
            return require($fileName);
        } else {
            throw new \Exception("文件路径错误", 1);
        }
    }

    public static function get()
    {
        func_get_args();   
    }

    public static function set()
    {
    
    }
}

