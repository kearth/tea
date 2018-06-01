<?php

namespace Tea\Core\Base;

final class Stream
{
    private static $class;

    private static $instance;

    private static $result;

    public final static function from(string $class){
        static::$class = Alias::get()[$class];
        if (static::$instance == null) {
            static::$instance = new static;
        }
        return static::$instance;
        throw new \Exception("不存在这个类");
    }

    public function __call($name, $args)
    {
        static::$result = forward_static_call_array([static::$class, $name], $args);
        return static::$instance;
    } 
}

