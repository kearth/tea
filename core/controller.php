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
        $ctr = Container::make($controller);
        $ctr->request  = $paramters;
        $ctr->$action();
        if (isset($ctr->response['type'])) {
            $returnValue = Container::make($ctr->response['type']);
            $returnValue->set($ctr->response);
            return $returnValue;
        }
        throw new \Exception('response is invailed');
    }
}

