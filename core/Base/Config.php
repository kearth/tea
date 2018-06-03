<?php

namespace Tea\Core\Base;

final class Config
{

    public static $appConfig = [];
    public static $routeConfig = [];

    public static function init(string $initPath)
    {
        static::loadOnly($initPath);
    }

    //public static function init()
    //{
        //static::$appConfig = static::load(APP_CONFIG);
        //static::$routeConfig = static::load(ROUTE_CONFIG);
    //}

    private static function loadOnly(string $fileName) : void
    {
        is_file($fileName) ? require($fileName) : throw new \Exception("文件不存在");
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
        $keyLayout = func_get_args();
        var_export($keyLayout);
    }

    public static function set()
    {
    
    }
}

