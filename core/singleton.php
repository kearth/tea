<?php
namespace Akf\Core;

abstract class Singleton
{
    private static $instance = [];
    private function __construct() 
    {
        $this->init();
    }

    final public static function getInstance()
    {
        $class = get_called_class();
    
        if (! isset(self::$instance[$class])) {
            self::$instance[$class] = new $class();
        }
        
        return self::$instance[$class];
    }

    final private function __clone()
    {
    
    }

    final private function __wakeup()
    {
    
    }

    final private function __sleep()
    {
    
    }
    
    protected function init()
    {
    
    }
    
}
