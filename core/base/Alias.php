<?php

namespace Tea\Core\Base;

class Alias
{
    private static $aliasList;

    public static function set($aliasList)
    {
       static::$aliasList = $aliasList;
    }

    public static function get()
    {
        return static::$aliasList;
    }
}
