<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Request;
use BaseStone\Core\Response;

class Dispatcher extends Base
{
    private $request;

    public function __construct()
    {
        $this->request  = Request::getInstance();
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
            $method = new $provider();
            $method->getAction();
            $method->output();
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

