<?php
namespace BaseStone\Core;

class Request extends Singleton
{
 
    private $info;
    private $provider;
    private $action;
    private $params;
    private $userAgent;
    private $protocal;
    private $uri;
    private $method;
    private $cookie;

    protected function init()
    {
        $this->info            = $_SERVER;
        $this->userAgent       = $_SERVER['HTTP_USER_AGENT'];
        $this->protocal        = $_SERVER['SERVER_PROTOCOL'];
        $this->uri             = $_SERVER['REQUEST_URI'];
        $this->method          = $_SERVER['REQUEST_METHOD'];
        $this->cookie          = $_COOKIE;
        $this->action          = $_REQUEST['rewrite'];
        $this->params          = $_REQUEST;
    }


    public function getUserAgent()
    {
        return $this->userAgent;
    }

    public function getInfo()
    {
        return $this->info;
    }

    //public function getProvider()
    //{
        //return $this->provider;
    //}

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

    //public function setProvider(string $provider)
    //{
        //$this->provider = $provider;   
    //}

    public function setAction(string $action)
    {
        $this->action = $action;
    }

    public function __call(string $name, array $paramters)
    {
        $type = substr($name, 0, 3);
        $property = lcfirst(substr($name, 3));
        if ('get' === $type && property_exists(get_class(), $property)) {
            return $this->$property;
        } elseif ('set' === $type && property_exists(get_class(), $property)) {
            $this->$property = $paramters[0];          
        }
    }   

}

