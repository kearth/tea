<?php

namespace Tea\Framework;

/**
 * 自动加载类
 */
class Autoload {

    /**
     * 自动加载映射
     */
    public static $autoloadMap = array(
        // 框架命名空间对应目录映射
        'Tea\Framework'   => 'framework',
        // App命名空间对应目录映射
        'App\Controllers' => 'app/controllers',
        'App\Lib'         => 'app/lib',
        'App\Models'      => 'app/models',
        'App\Utils'       => 'app/utils'
    );

    /**
     * 已加载类
     */
    public static $hasLoad = array();

    /**
     * 初始化
     */
    public static function init() : void {
        // 注册自动加载机制
        spl_autoload_register(function($className){
            // 已经加载的不需要重复加载
            if (isset(self::$hasLoad[$className])) {
                return false;
            }
            $nameList = explode('\\', $className);
            $class = array_pop($nameList);
            $namespace = implode('\\', $nameList);
            if (isset(self::$autoloadMap[$namespace])) {
                // 有映射的按照映射加载
                $realClass = ROOT_PATH . DIRECTORY_SEPARATOR . self::$autoloadMap[$namespace] . DIRECTORY_SEPARATOR . $class . '.php';
            } else {
                // 没有映射按默认规则加载, 目录小写, 类名驼峰
                $realClass = ROOT_PATH . DIRECTORY_SEPARATOR . strtolower($namespace) . DIRECTORY_SEPARATOR . $class . '.php';
            }
            if (is_file($realClass)) {
                include($realClass);
            } else {
                throw new \Exception($className . " 类文件不存在", CODE_FRAMEWORK);
            }
            // 加载过的类添加到对应map
            self::$hasLoad[$className] = $realClass;
        });
    }

}

