<?php
namespace Akf\Core;

class Dispatcher extends Base
{
    private $request;

    public function __construct()
    {
        $this->request  = Request::getInstance();;
    }
    
    public function run()
    {
        $this->preDispatch();
        $this->dispatch();
        $this->postDispatch();   
    }

    public function dispatch()
    {
        $provider = $this->request->getProvider();
        if (class_exists($provider)) {
            $method = new $provider($this->request, Response::getInstance());
            $method->getAction();
        } else {
            echo "请求不存在";
        }
    }

    public function preDispatch()
    {
    
    }

    public function postDispatch()
    {
    
    }

}

