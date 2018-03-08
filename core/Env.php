<?php

namespace Tea\Core;

class Env
{
    private static $envList = [];

    public static function init(array $envList)
    {
        foreach ($envList as $name => $value) {
            self::set($name, $value);
        }
    }

    public static function get(string $envName)
    {   
        if (isset(self::$envList[$envName])) {
            return self::$envList[$envName];
        }
    }

    public static function set(string $envName, $value)
    {
        if (!defined($envName)) {
            define($envName, $value);
            self::$envList[$envName] = $value;
        }
    }
}
