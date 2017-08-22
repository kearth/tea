<?php
namespace Akf\Core;

class Dispatcher extends Base
{
    private $request;

    public function __construct()
    {
        $this->request  = Request::getInstance();;
    }
    
    public static function run()
    {
        $self = self::getInstance();
        $self->preDispatch();
        $self->dispatch();
        $self->postDispatch();   
    }

    public function dispatch()
    {
        $provider = $this->request->getProvider();
        if (class_exists($provider)) {
            $method = new $provider($this->request, Response::getInstance());
            $method->getAction();
            //$method->output();
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

