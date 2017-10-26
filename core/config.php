<?php
namespace Akf\Core;

class Config
{
    private static $config = [];

    /**
     *  加载配置文件 load
     *
     *  @param $configPath 配置文件路径
     *  @return bool
     */   
    public static function loadDefaultCfg(string $configPath) : bool
    {
        if (file_exists($configPath)) {
            self::$config = include $configPath;
            define('ENV', self::get('env'));
            return true;
        }
        return false;
    }

    public static function getComponents()
    {
        return self::get('components');
    }

    /**
     * @name get 获取配置
     * @param $name 配置名
     * @return mixed
     */
    public static function get($name)
    {
        if (array_key_exists($name, self::$config)) {
            if (is_string(self::$config[$name]) &&  file_exists(self::$config[$name])) {
                return include self::$config[$name];
            } else {
                return self::$config[$name];
            }
        }
    }
}

