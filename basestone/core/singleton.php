<?php
namespace BaseStone\Core;

class Singleton
{
    private static $instance = null;

    private function __construct()
    {
    
    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new static();
        }
        return self::$instance;
    }

}

