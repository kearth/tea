<?php
namespace BaseStone\Core;

interface IReflection
{
    public static function execute($class, $method, $paramters, $isStatic);
    public static function getObject($class, $paramters);
    public static function doMethod($class, $method, $paramters, $isStatic);
    public static function bindExtend($sonClass, $parentClass);
    public static function bindImplements($class, $interface);
}
