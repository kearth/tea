<?php

namespace Tea\Kernel;


class Autoload
{
    public static $hasRegister = [];

    public static function initAlias() : void
    {
        $alias = Config::get(CONFIG_KEY_AUTOLOAD, AUTOLOAD_ALIAS);
        if (!empty($alias)) {
            var_export($alias);
            foreach ($alias as $class => $name) {
                class_alias($class, $name, true);
            }
        }
    }

    public static function register() : void
    {
        spl_autoload_register(function($class) {
            if (isset(self::$hasRegister[$class])) return;
            $classArr = explode("\\", $class);
            if (!empty($classArr)) {
                if ($classArr[0] === Config::get(CONFIG_KEY_AUTOLOAD, FRAMEWORK_NAME)) {
                    array_shift($classArr);
                    $classPathRoot      = $root . DIRECTORY_SEPARATOR . implode(DIRECTORY_SEPARATOR, $classArr);
                    $classPath          = $classPathRoot . ".php";

                    if (is_file($classPath)) {
                        require($classPath);
                        self::$hasRegister[$class] = $classPath;
                    }

                } else {
                    //TODO 外部类
                }
            }
        });
    }
}
