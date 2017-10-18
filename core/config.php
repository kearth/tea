<?php
namespace Akf\Core;

class Config
{
    private static $config = [];
    
    public static function load($configPath) : bool
    {
        if (file_exists($configPath)) {
            self::$config = include $configPath;
            return true;
        }
        return false;
    }

    public static function getEnv()
    {
        return self::get('env');
    }

    public static function getComponents()
    {
        return self::get('components');
    }

    public static function getComponentCfg(string $cfg) : array
    {
        if (file_exists($cfg)) {
            return include $cfg;
        }
        return [];
    }


    /**
     * @name get 获取配置
     * @param $name 配置名
     * @return mixed
     */
    public static function get($name)
    {
        if (array_key_exists($name, self::$config)) {
            return self::$config[$name];
        }
    }

    /**
     * @name set 设置配置
     * @param $name 配置名
     * @param $value 配置值
     * @return void
     */
    public static function set($name,$value)
    {
        self::$config[$name] = $value;
    }
}

