<?php
namespace BaseStone\Core;

class AutoReflection implements IReflection
{
    public static function execute($class, $method, $paramters, $isStatic = false)
    {
    
    }

    public static function getObject($class, $paramters = [])
    {
        if (class_exists($class)) {
            $reflectionClass = new \ReflectionClass($class);
            return $reflectionClass->newInstance($paramters);
        } 
        throw new \ReflectionException("class not founed");
    }

    public static function doMethod($class, $method, $paramters, $isStatic = false)
    {
    
    }

    public static function bindExtend($sonClass, $parentClass)
    {
    
    }

    public static function bindImplements($class, $interface)
    {
    
    }
    
}
