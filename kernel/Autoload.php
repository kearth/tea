<?php
namespace Akf\Kernel;

class Autoload
{
    public static $nsHead      = "Akf";
    public static $hasRegister = [];
    public static function register(string $root) : void
    {
        spl_autoload_register(function($class) use ($root){
            if (isset(self::$hasRegister[$class])) return;
            $classArr = explode("\\", $class);
            if (!empty($classArr)) {
                if ($classArr[0] === self::$nsHead) {
                    array_shift($classArr);
                    $classPathRoot      = $root . DIRECTORY_SEPARATOR . implode(DIRECTORY_SEPARATOR, $classArr);
                    $classPath          = $classPathRoot . ".php";
                    $classPathAbstract  = $classPathRoot . "Abstract.php";
                    $classPathInterface = $classPathRoot . "Interface.php";

                    if (file_exists($classPath)) {
                        require $classPath;
                        self::$hasRegister[$class] = $classPath;
                    } else if (file_exists($classPathAbstract)) {
                        require $classPathAbstract;
                        self::$hasRegister[$class] = $classPathAbstract;
                    } else if (file_exists($classPathInterface)) {
                        require $classPathInterface;
                        self::$hasRegister[$class] = $classPathInterface;
                    }   

                } else {
                    //TODO 外部类
                }
            }
        });
    }
}

