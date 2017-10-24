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
    public static function load(string $configPath) : bool
    {
        if (file_exists($configPath)) {
            self::$config = include $configPath;
            return true;
        }
        return false;
    }
    
    public static function getAliasCfg()
    {
        return include self::get('alias');
    }

    public static function getBindCfg()
    {
        include self::get('bind');
    }

    public static function getEnv()
    {
        return self::get('env');
    }

    public static function getComponents()
    {
        return self::get('components');
    }

    /**
     *  获取组件配置 getComponents
     *
     *  @param $cfg 配置文件路径
     *  @return array
     */
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

