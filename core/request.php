<?php

namespace Akf\Core;

final class Request
{
    const HTTP_METHOD_GET  = 'GET';
    const HTTP_METHOD_POST = 'POST';
    
    private $httpAccept    = '';
    private $httpUserAgent = '';
    private $requestMethod = '';
    private $request       = [];

    public function __construct()
    {
        $this->httpAccept    = $_SERVER['HTTP_ACCEPT'];
        $this->httpUserAgent = $_SERVER['HTTP_USER_AGENT'];
        $this->requestMethod = $_SERVER['REQUEST_METHOD'];
        $this->request       = $_REQUEST;
        $this->check();
    }

    private function check()
    {
        if (
            empty($this->httpAccept)    || 
            empty($this->httpUserAgent) ||
            empty($this->requestMethod) ||
            empty($this->request)     
        ) {
            throw \Exception('Request init failed');
        }
    }

    public function getAccept()
    {
    
    }

    public function getUserAgent()
    {
    
    }

    public function getRequestMethod() : string
    {
        return (self::HTTP_METHOD_GET === $this->requestMethod) ? self::HTTP_METHOD_GET : self::HTTP_METHOD_POST;
    }

    public function getRequest()
    {
    
    }
}
