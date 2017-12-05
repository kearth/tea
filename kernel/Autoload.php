<?php
namespace Akf\Kernel;

class Autoload
{
    public static function register(string $root) : void
    {
        spl_autoload_register(function($class) use ($root){
            echo $root;       
            echo $class;
            //if (file_exists(CONTROLLER . $class . '.php')) {
                //$classFile = CONTROLLER . $class . '.php';
            //} elseif (file_exists(MODEL . $class . '.php')) {
                //$classFile = MODEL . $class . '.php';
            //} else {
                //$classFile = ROOT_PATH . DIRECTORY_SEPARATOR . substr(str_replace('\\', '/', $class), 4) . '.php';
            //}

            //if (in_array($classFile, self::$hasRequiredFile)) {
                //return;
            //}

            //if (file_exists($classFile)) {
                ////echo $classFile . "<br>";
                //require $classFile;
                //self::$hasRequiredFile[] = $classFile;
            //}

        });
    }
}

