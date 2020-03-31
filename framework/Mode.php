<?php

namespace Tea\Core;

class Mode 
{
    public static $mode;

    public final static function detectWhichMode()
    {
        $sapiName = php_sapi_name();
        switch ($sapiName){
            case "cli":
                static::$mode = "cli";
                break;
            case "fpm-fcgi":
                static::$mode = "fpmFcgi";
                break;
            default:
                throw new \Exception("暂不支持这种模式");
        }
    }

    public final static function initTheMode()
    {
        $method = static::$mode . 'Mode';
        forward_static_call_array([__class__, $method], []);
    }

    private static function cliMode()
    {
    
    }

    private static function fpmFcgiMode()
    {
        echo "start fpmFcgi";
    }
}
