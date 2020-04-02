<?php

namespace Tea\Framework;

/**
 * 配置类
 */
class Config {

    /**
     * 全部配置, 根据配置文件名分key存储
     */
    public static $configMap = array();

    /**
     * 初始化
     */
    public static function init() : void {
        if (!is_dir(CONF_ROOT)) {
            throw new \Exception("配置路径不存在", CODE_FRAMEWORK);
        }
        // 遍历配置目录
        $configFileList = scandir(CONF_ROOT);
        foreach($configFileList as $file) {
            if ($file === '.' || $file === '..') {
                continue;
            }
            $conf = str_replace(strrchr($file, "."), "", $file);
            static::$configMap[$conf] = static::load(CONF_ROOT . $file);
        }
    }

    /**
     * 加载配置
     */
    public static function load(string $fileName) : array {
        if (is_file($fileName)) {
            return include($fileName);
        } 
        throw new \Exception("配置文件不存在", CODE_FRAMEWORK);
    }

    /**
     * 获取配置
     */
    public static function get(string $name, $config = DEFAULT_CONF_KEY) {
        if (isset(static::$configMap[$config]) && isset(static::$configMap[$config][$name])) {
           return static::$configMap[$config][$name]; 
        }
        throw new \Exception("配置不存在", CODE_FRAMEWORK);
    }

    /**
     * 获得一个key下全部配置
     */
    public static function getAll($config = DEFAULT_CONF_KEY) {
        if (isset(static::$configMap[$config])) {
           return static::$configMap[$config];
        }
        throw new \Exception("配置不存在", CODE_FRAMEWORK);
    }

    /**
     * 设置配置 
     */
    public static function set(string $name, $value, $config = DEFAULT_CONF_KEY) {
        if (isset(static::$configMap[$config])) {
            static::$configMap[$config][$name] = $value;
        } else {
            static::$configMap[$config] = array($name => $value);
        }
    }

    /**
     * 通过类名初始化配置
     */
    public static function initConf(string $class) {
        $realClass = strtolower(basename(str_replace('\\', '/', $class)));
        return static::getAll($realClass);   
    }

}

