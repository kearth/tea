<?php
namespace Akf\Core;

abstract class Service
{
    public function __construct(Request $request, Response $response)
    {
    }

    public function indexAction()
    {
    
    }

    public function beforeService(Request $request, Response $response)
    {
    
    }

    public function afterService(Request $request, Response $response)
    {
    
    }

    public function __call(string $name, array $args)
    {
        if (substr($name, -6) == 'Action') {
            $this->beforeService($request, $response);
            $this->$name($args);
            $this->afterService($request,$response);    
        }
    }
}
