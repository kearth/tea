<?php

namespace Akf\Core\BaseSource;

final class Stream extends BaseSource
{
    private $uri        = null;
    private $controller = null;
    private $action     = null;
    private $param      = [];

    /**
     *  请求
     */
    private $request = null;

    /**
     *  响应
     */
    private $response = null;

    public function __construct(Request $request)
    {
        $this->request  = $request;
    }

    public function getUri() : string
    {
        if (is_null($this->uri)) {
            $this->uri = $this->request->getUri();
        }
        return $this->uri;
    }   
    
    public function getController() : string
    {
        return $this->controller;
    }

    public function getAction() : string 
    {
        return $this->action;
    }

    public function getParam() : array
    {
        return $this->request->getParam();
    }
    
    public function setUri(string $uri)
    {
        $this->uri = $uri;
    }

    public function setController(string $controller)
    {
        $this->controller = $controller;
    }

    public function setAction(string $action)
    {
        $this->action = $action;
    }

    public function getResponse() : Response
    {
        return $this->response;
    }

    public function setResponse(Response $value) 
    {
        $this->response = $value;
    }

    public function out()
    {
        exit;
    }
}

