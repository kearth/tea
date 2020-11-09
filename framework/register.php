<?php 
namespace Tea\Framework;

class Register {

    private static array $instanceList = array();

    public static function add(string $name, object $instance) : bool{
        if(!isset(self::$instanceList[$name])){
            self::$instanceList[$name]  = $instance;
            return true;
        }
        return false;
    }

    public static function getInstance(string $name) : ?object {
        return self::$instanceList[$name] ?? null; 
    }
}
