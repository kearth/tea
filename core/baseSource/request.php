<?php

namespace Akf\Core\BaseSource;

final class Request
{
    const HTTP_METHOD_GET  = 'GET';
    const HTTP_METHOD_POST = 'POST';
    
    private $httpAccept    = '';
    private $httpUserAgent = '';
    private $requestMethod = '';
    private $uri           = '';
    private $param         = [];

    public function __construct()
    {
        $this->httpAccept    = $_SERVER['HTTP_ACCEPT'];
        $this->httpUserAgent = $_SERVER['HTTP_USER_AGENT'];
        $this->requestMethod = $_SERVER['REQUEST_METHOD'];
        $this->uri           = $_REQUEST['rewrite'];
        $this->param         = array_diff_key($_REQUEST, ['rewrite' => []]); 
        $this->check();
    }

    private function check()
    {
        if (
            empty($this->httpAccept)    || 
            empty($this->httpUserAgent) ||
            empty($this->requestMethod) ||
            empty($this->uri)     
        ) {
            throw \Exception('Request init failed');
        }
    }

    public function getAccept()
    {
        return $this->httpAccept;  
    }

    public function getUserAgent()
    {
        return $this->httpUserAgent;
    }

    public function getRequestMethod() : string
    {
        return (self::HTTP_METHOD_GET === $this->requestMethod) ? self::HTTP_METHOD_GET : self::HTTP_METHOD_POST;
    }

    public function getUri()
    {
        return $this->uri;
    }

    public function getParam()
    {
        return $this->param;
    }
}
