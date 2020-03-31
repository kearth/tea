<?php

namespace Tea\Core;

/**
 *
 */
class Container
{
    private static $bindList = [];

    private const TYPE_CLASS = "class";
    private const TYPE_CLOSURE = "closure";
    private const TYPE_INSTANCE = "instance";

    public static function get()
    {
        $args = func_get_args();
        $num  = func_num_args();

        if ($num == 1) {
            $key = array_shift($args);
            if (isset(self::$bindList[$key])) {
                $bind = self::$bindList[$key];
                $type = ucfirst($bind[0]);
                $value = $bind[1];
                if (!isset($bind[2])) {
                    self::$bindList[$key][2] = call_user_func([__CLASS__, "get" . $type], $value, []);
                }
                return self::$bindList[$key][2];
            }
        }

        if ($num == 2) {
            $key = array_shift($args);
            if (isset(self::$bindList[$key])) {
                $bind = self::$bindList[$key];
                $type = ucfirst($bind[0]);
                $value = $bind[1];
                if (!isset($bind[2]) || $args[0] === true) {
                    $args[0] = is_array($args[0]) ? $args[0] : [];
                    self::$bindList[$key][2] = call_user_func([__CLASS__, "get" . $type], $value, $args[0]);            
                }
                return self::$bindList[$key][2];
            }
        }

        throw new \Exception("error");
    }

    private static function getClass(string $className, array $arr)
    {
        $class = new \ReflectionClass($className);
        return $class->newInstanceArgs($arr);
    }

    private static function getClosure(\Closure $closure, array $arr)
    {
        if (!empty($arr)) {
            return call_user_func_array($closure, $arr);
        }
    }

    private static function getInstance(string $value, array $arr)
    {
        return $value;
    }

    public static function bind()
    {
        $args = func_get_args();
        $num  = func_num_args();

        switch ($num) {
            case 1 : 
                if (is_array($args[0])) {
                    foreach ($args[0] as $key => $value) {
                        self::bind($key, $value);
                    }
                    break;
                }
                //TODO Exception
                throw new \Exception("没绑定上");
            case 2 : 
                if (is_string($args[0])) {
                    $key = array_shift($args);
                    if ($args[0] instanceof \Closure) {
                        self::bindClosure($key, $args[0]);
                        break;
                    }

                    if (is_object($args[0])) {
                        self::bindInstance($key, $args[0]);
                        break;
                    }

                    if (class_exists($args[0])) {
                        self::bindClass($key, $args[0]);
                        break;
                    }
    
                }
                throw new \Exception("没绑定上");
            default :
                throw new \Exception("没绑定上");
        }
    }

    private static function bindClass(string $key, string $class)
    {
        self::$bindList[$key] = [
            self::TYPE_CLASS,
            $class
        ];
    }

    private static function bindClosure(string $key, \Closure $closure)
    {
        self::$bindList[$key] = [
            self::TYPE_CLOSURE,
            $closure
        ];
    }

    private static function bindInstance(string $key, object $object)
    {
        self::$bindList[$key] = [
            self::TYPE_INSTANCE,
            $object
        ];
    }

}