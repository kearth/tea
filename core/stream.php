<?php
namespace Akf\Core;

class Stream 
{
    private static $request  = [];
    private static $response = [];

    public function __construct(array $content)
    {
        self::$request['uri'] = $content['rewrite'];
    }

    public function getRequest($name = '')
    {
        if ('' !== $name) {
            return self::$request[$name];
        }  else {
            return self::$request;
        }
    }   

    public function setRequest(string $name, $value)
    {
        self::$request[$name] = $value;
    }


    public function getResponse($name = '')
    {
    
    }
}
