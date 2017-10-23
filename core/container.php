<?php

namespace Akf\Core;

class Container
{
    private static $instance = [];

    public static function bind(string $class, \Closure $closure)
    {
        self::$instance[$class] = $closure;
    }

    public function make(string $class, $paramters = [])
    {
        $closure = self::$instance[$class];
        return $closure($paramters);
    }
}

