<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Response;

class Dispatcher
{
    private static $instance = null;
    private $response;

    private function __construct()
    {
        $this->response = Response::getInstance();
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

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    public function dispatch($request)
    {
        $action = "Application\\".str_replace('/','\\',$request->action);
        if (class_exists($action)) {
            $method = new $action();
            $method->getAction();
        } else {
            echo "请求不存在";
        }
        $this->response();
    }

    public function response()
    {
        $response = $this->response->getResponse();
        var_dump($response);
    }

}

