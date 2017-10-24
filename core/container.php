<?php

namespace Akf\Core;

class Container
{
    private static $instance = [];

    public static function bind(string $class, \Closure $closure)
    {
        self::$instance[$class] = $closure;
    }

    public static function make(string $class, $paramters = [])
    {
        if (array_key_exists($class, self::$instance)) {
            $closure = self::$instance[$class];
            return $closure($paramters);
        }
        throw new \Exception('no this class :' . $class);
    }
}

