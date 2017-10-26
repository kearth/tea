<?php

namespace Akf\Core;

class Container
{
    private static $instance = [];

    public static function init(array $cfg)
    {
        self::$instance = $cfg;   
    }

    public static function bind(string $class, \Closure $closure)
    {
        self::$instance[$class] = $closure;
    }

    public static function make(string $class, $paramters = [])
    {

        if (array_key_exists($class, self::$instance)) {
            $closure = self::$instance[$class];
            return $closure($paramters);
        } elseif (class_exists($class)) {
            self::bind($class, function ($param) use ($class) {
                return new $class($param);
            });
            return self::make($class, $paramters);
        }
        throw new BaseException('this class has no binded :' . $class);
    }

}

