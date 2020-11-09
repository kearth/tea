<?php

namespace Tea\Framework;

/**
 * 自动加载类
 */
class Autoload {

    /**
     * 已加载类
     */
    public static array $hasLoad = array();

    /**
     * 初始化
     */
    public static function init(string $rootPath) : void {
        // 注册自动加载机制
        spl_autoload_register(function($className) use($rootPath){
            // 已经加载的不需要重复加载
            if (isset(self::$hasLoad[$className])) {
                return false;
            }
            $nameList = explode('\\', $className);
            $classFile = self::getUnderScoreClassFile(array_pop($nameList));
            $namespace = self::getNameSpace($nameList);
            $realClass = self::getRealClass(dirname($rootPath), strtolower($namespace), $classFile);
            if (is_file($realClass)) {
                require($realClass);
                // 加载过的类添加到对应map
                self::$hasLoad[$className] = $realClass;
            }
        });
    }

    private static function getRealClass(string $dir, string $namespace, string $classFile){
        return implode(DIRECTORY_SEPARATOR, array($dir, $namespace, $classFile));
    }
    private static function getNameSpace(array $nameList){
        return implode(DIRECTORY_SEPARATOR, $nameList);
    }

    private static function getUnderScoreClassFile(string $str) : string {
        $newList = array();
        $charList = str_split(lcfirst($str));
        foreach($charList as $char) {
            $ord = ord($char);
            if(65 <= $ord && $ord <= 90){
                $newList[] = "_" . chr($ord + 32); 
                continue;
            } 
            $newList[] = $char;
        }
        $newList[] = ".php";
        return implode("", $newList);
    }

}

