<?php

namespace Tea\Core;

class Request
{
    private static $name = "Request";

    const KEY_SERVER = "SERVER";
    const KEY_USER = "USER";
    const KEY_DISPATCH = "DISPATCH";

    private static $request = [];

    public static function init()
    {
        self::$request["SERVER"] = $_SERVER;
    }

    public static function get()
    {
        return self::$request;  
    }
    
    public static function set()
    {
    
    }

}
