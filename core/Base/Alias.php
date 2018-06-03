<?php

namespace Tea\Core\Base;

class Alias
{
    private static $aliasList;

    public static function set(array $aliasList) : void
    { 
       static::$aliasList = $aliasList;
    }

    public static function get() : array 
    {
        return static::$aliasList;
    }
}
