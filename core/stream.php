<?php
namespace Akf\Core;

class Stream 
{
    /**
     *  请求
     */
    private static $request  = [];

    /**
     *  响应
     */
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
        if ('' !== $name) {
            return self::$response[$name];
        }  else {
            return self::$response;
        }
    }

    public function setResponse(string $name, $value) 
    {
        self::$response[$name] = $value;
    }

    public function flowAway()
    {
        exit;
    }
}
