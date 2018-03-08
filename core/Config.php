<?php

namespace Tea\Core;

class Config
{
    private static $configList = [];

    private static $keyList = [];

    /**
     * 初始化配置加载
     * @param string $fileName 文件路径
     */
    public static function init(string $initPath) : void
    {
        $initConfig = self::load($initPath);
        foreach ($initConfig as $key => $value) {
            self::$configList[$key] = $value;
        }
    }

    private static function load(string $fileName)
    {
        if (is_file($fileName)) {
            return require($fileName);
        }
        throw new \Exception("文件路径错误", 1);
    }

    public static function getConfigAll()
    {
        return self::$configList;
    }

    public static function setConfig(string $key, $value)
    {
        self::$configList[$key] = $value;
    }

    public static function getConfig(string $key)
    {
        if (array_key_exists($key, self::$configList)) {
            return self::$configList[$key];
        }
    }
}
