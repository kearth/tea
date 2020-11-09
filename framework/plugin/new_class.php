<?php 

namespace Tea\Framework\Plugin;

class NewClass extends Plugin {

    private array $aliasList = array();

    public static function make(string $class) : ?object {
        if(class_exists($class)) {
            if (isset(self::getInstance()->aliasList[$class])) {
                $newClass = self::getInstance()->aliasList[$class];
                return new $newClass(); 
            }
            return new $class(); 
        }
        return null;
    }

    public static function alias(string $class, string $alias) : bool {
        if(class_exists($class) && class_exists($alias)) {
            self::getInstance()->aliasList[$class] = $alias;
            return true;
        }
        return false;
    }
}
