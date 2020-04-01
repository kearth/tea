<?php

namespace Tea\Framework;

/**
 * 配置类
 */
class Config {

    public static $confRoot = ROOT_PATH . '/conf/';

    public static $configMap = array();

    public static function init() {
        if (!is_dir(static::$confRoot)) {
            throw new \Exception("配置路径不存在", 1);
        }
        $configFileList = scandir(static::$confRoot);
        foreach($configFileList as $file) {
            if ($file === '.' || $file === '..') {
                continue;
            }
            $conf = str_replace(strrchr($file, "."), "", $file);
            static::$configMap[$conf] = static::load(static::$confRoot . $file);
        }
    }

    public static function load(string $fileName) {
        if (is_file($fileName)) {
            return include($fileName);
        } 
        throw new \Exception("配置文件不存在", 1);
    }

    public static function get($name, $config = 'init') {
        if (isset(static::$configMap[$config]) && isset(static::$configMap[$config][$name])) {
           return static::$configMap[$config][$name]; 
        }
        throw new \Exception("配置不存在", 1);
    }

    public static function set($name, $value, $config = 'init') {
        if (isset(static::$configMap[$config])) {
            static::$configMap[$config][$name] = $value;
        } else {
            static::$configMap[$config] = array($name => $value);
        }
    }
}

