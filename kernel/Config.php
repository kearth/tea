<?php

namespace Tea\Kernel;

class Config
{
    private static $configList = [];

    private static $keyList = [];

    /**
     * 初始化配置加载
     * @param string $fileName 文件路径
     */
    public static function initLoad(string $fileName) : void
    {
        $initConfig = self::load($fileName);
        foreach ($initConfig as $rootKey => $value) {

            if (is_string($value) && is_file($value)) {
                $value = self::load($value);
            }

            if(is_array($value)){
                self::$configList[$rootKey] = $value;
                self::$keyList[] = $rootKey;
            }
        }
    }

    private static function load(string $fileName) : array
    {
        if (is_file($fileName)) {
            return require($fileName);
        }
        throw new \Exception("文件路径错误", 1);
    }

    public static function set() : void
    {
        $numargs = func_num_args();
        $args = func_get_args();

        switch ($numargs) {
            case 3:
                $configRootKey = $args[0];
                $configKey = $args[1];
                $configValue = $args[2];
                break;
            default:
                throw new \Exception("config set has something wrong");
                break;
        }

        if (in_array($configRootKey, self::$keyList)) {
            self::setConfig($configRootKey, $configKey, $configValue);
        }
    }

    public static function get()
    {
        $numargs = func_num_args();
        $args = func_get_args();

        switch ($numargs) {
            case 2:
                $configRootKey = $args[0];
                $configKey = $args[1];
                break;
            default:
                throw new \Exception("Error Processing Request", 1);
                break;
        }

        return self::getConfig($configRootKey, $configKey);
    }


    private static function setConfig(string $rootKey, string $key, string $value) : void
    {
        self::$configList[$rootKey][$key] = $value;
    }

    private static function getConfig(string $rootKey, string $key)
    {
        if (array_key_exists($rootKey, self::$configList)) {
            if (array_key_exists($key, self::$configList[$rootKey])) {
                return self::$configList[$rootKey][$key];
            }
            return self::$configList[$rootKey];
        }
    }

    private static function getFileType(string $fileName) : string
    {
        $fileInfo = pathinfo($fileName);
        return $fileInfo["extension"];
    }
}
