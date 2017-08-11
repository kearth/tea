<?php
namespace BaseStone\Core;

class Request extends Core
{
 
    private static $instance = null;
    private $info;
    private $provider;
    private $action;
    private $params;
    private $userAgent;
    private $protocal;
    private $uri;
    private $method;
    private $cookie;

    private function __construct()
    {
        $this->info            = $_SERVER;
        $this->userAgent       = $_SERVER['HTTP_USER_AGENT'];
        $this->protocal        = $_SERVER['SERVER_PROTOCOL'];
        $this->uri             = $_SERVER['REQUEST_URI'];
        $this->method          = $_SERVER['REQUEST_METHOD'];
        $this->cookie          = $_COOKIE;
        $this->action          = $_REQUEST['rewrite'];
        unset($_REQUEST['rewrite']);
        $this->params          = $_REQUEST;
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

    public function getUserAgent()
    {
        return $this->userAgent;
    }

    public function getInfo()
    {
        return $this->info;
    }

    public function getProvider()
    {
        return $this->provider;
    }

    public function getAction()
    {
        return $this->action;
    }

    public function getParams()
    {
        return $this->params;
    }

    public function getProtocal()
    {
        return $this->protocal;
    }

    public function getUri()
    {
        return $this->uri;
    }

    public function getMethod()
    {
        return $this->method;   
    }

    public function getCookie()
    {
        return $this->cookie;
    }

    public function setProvider(string $provider)
    {
        $this->provider = $provider;   
    }

    public function setAction(string $action)
    {
        $this->action = $action;
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }
       
}

