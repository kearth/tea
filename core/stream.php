<?php

namespace Akf\Core;

class Stream extends BaseSource
{
    /**
     *  请求
     */
    private static $request  = [];

    /**
     *  响应
     */
    private static $response;

    public function __construct(array $content)
    {
        self::$request['uri']   = $content['rewrite'];
        unset($content['rewrite']);
        self::$request['param'] = $content;
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


    public function getResponse()
    {
        return self::$response;
    }

    public function setResponse(Response $value) 
    {
        self::$response = $value;
    }

    public function out()
    {
        exit;
    }
}

