<?php

namespace Tea\Core\Base;

final class Stream
{
    private static $class;

    private static $instance;

    private static $params;

    public static function from(string $class)
    {
        if (array_key_exists($class, Alias::get())){
            self::$class = Alias::get()[$class];
            if (self::$instance == null) {
                self::$instance = new static;
            }
            return self::$instance;
        } 
        throw new \Exception("不存在这个类");
    }

    public function __call($name, $args)
    {
        self::$params = forward_static_call_array([self::$class, $name], $args);
        return $this;
    }
}

