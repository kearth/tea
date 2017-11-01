<?php

namespace Akf\Core\Kernel;

class Container
{
    private static $instance  = [];
    private static $singleton = [];

    public static function init(array $cfg)
    {
        self::$instance = $cfg;   
    }

    public static function initSingleton(array $cfg)
    {
        self::$singleton = $cfg;
    }

    public static function bind(string $class, \Closure $closure)
    {
        self::$instance[$class] = $closure;
    }

    public static function bindSingle(string $class, \Closure $closure)
    {
        self::$singleton[$class] = $closure;
    }

    public static function singleton(string $class)
    {
        if (array_key_exists($class, self::$singleton)) {
            $singleton = self::$singleton[$class];
            if (is_callable($singleton)) {
                $paramters = func_get_args();
                array_shift($paramters);
                self::$singleton[$class] = call_user_func_array($singleton, $paramters);
                return self::$singleton[$class];
            } elseif (is_object($singleton)) {
                return $singleton;
            } else {
                throw new BaseException('this singleton class bind error:' . $class);
            }
        }   
        throw new BaseException('no this singleton class:' . $class);
    }

    public static function make(string $class)
    {
        $paramters = func_get_args();
        array_shift($paramters);
        if (array_key_exists($class, self::$instance)) {
            $closure = self::$instance[$class];
            return call_user_func_array($closure, $paramters);
        } elseif (class_exists($class)) {
            self::bind($class, function ($param) use ($class) {
                return new $class($param);
            });
            return self::make($class, $paramters);
        }
        throw new BaseException('this class has no binded :' . $class);
    }

}

