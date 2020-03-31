<?php

namespace Tea\Framework;

class Autoload {

    public static $autoloadMap = array(
        'Tea\Framework' => 'framework'
    );

    public static $hasLoad = array();

    public static function init() {
        spl_autoload_register(function($className){
            if (isset(self::$hasLoad[$className])) {
                return false;
            }
            $nameList = explode('\\', $className);
            $class = array_pop($nameList);
            $namespace = implode('\\', $nameList);
            if (isset(self::$autoloadMap[$namespace])) {
                $realClass = ROOT_PATH . DIRECTORY_SEPARATOR . self::$autoloadMap[$namespace] . DIRECTORY_SEPARATOR . $class . '.php';
            } else {
                echo $className;
                $realClass = ROOT_PATH . DIRECTORY_SEPARATOR . strtolower($namespace) . DIRECTORY_SEPARATOR . $class . '.php';
            }
            require $realClass;
            self::$hasLoad[$className] = $realClass;
        });
    }

}

