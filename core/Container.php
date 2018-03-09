<?php

namespace Tea\Core;

/**
 *
 */
class Container
{
    //类定义列表
    private static $classList = [];

    public static function get(string $class, $param = null)
    {
        if (isset(self::$classList[$class])) {
            $closure = self::$classList[$class];
            return is_null($param) ? $closure() : $closure($param);
        }
    }

    public static function set(string $class, \Closure $closure)
    {
        self::$classList[$class] = $closure;
    }
}
