<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Request;
use BaseStone\Core\Response;

class Dispatcher extends Base
{
    private $request;
    private $response;

    public function __construct()
    {
        $this->request  = Request::getInstance();
        $this->response = Response::getInstance();
    }
    
    public function run()
    {
        $this->preDispatch();
        $this->dispatch();
        $this->postDispatch();   
    }

    public function dispatch()
    {
        $action = "Application\\".str_replace('/','\\',$this->request->action);
        if (class_exists($action)) {
            $method = new $action();
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

