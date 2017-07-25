<?php
namespace BaseStone\Core;

class Request
{
    private static $instance = null;
    const DISPLAY_ARRAY = 0;
    const DISPLAY_OBJECT = 1;

    private function __construct()
    {

    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    private function package() 
    {
        $package = [];
        $package['HTTP_USER_AGENT'] = $_SERVER['HTTP_USER_AGENT'];
        $package['SERVER_PROTOCOL'] = $_SERVER['SERVER_PROTOCOL'];
        $package['REQUEST_URI']     = $_SERVER['REQUEST_URI'];
        $package['REQUEST_METHOD']  = $_SERVER['REQUEST_METHOD'];
        $package['CONTENT_LENGTH']  = $_SERVER['CONTENT_LENGTH'];
        $package['CONTENT_TYPE']    = $_SERVER['CONTENT_TYPE'];
        $package['COOKIE']          = $_COOKIE;
        $package['GET']             = $_GET;
        $package['POST']            = $_POST;
        $package['ENV']             = $_ENV;
        $package['ACTION']          = $_REQUEST['rewrite'];
        unset($_REQUEST['rewrite']);
        $package['PARAMS']          = $_REQUEST;
        return $package;
    }

    private function toObj($arr)
    {
        $obj = new \StdClass();

        foreach ($arr as $key => $value) {
            $obj->$key = $value;    
        }      

        return $obj;
    }


    public function getRequest($back_type = self::DISPLAY_OBJECT)
    {
        if (self::DISPLAY_OBJECT === $back_type) {
            return $this->toObj($this->package());    
        } else {
            return $this->package();
        }
    }
       
}

