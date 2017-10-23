<?php
namespace Akf\Core;

class Controller
{
    public function __toString()
    {   
        return get_class($this);
    }

    public function __invoke(string $controller, string $action, array $paramters)
    {
        $ctr = new $controller;
        $ctr->request  = $paramters;
        $ctr->response = '';
        $ctr->$action();
        return $ctr->response;
    }
}

