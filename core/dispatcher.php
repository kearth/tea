<?php

namespace Akf\Core;

class Dispatcher extends BaseComponent
{
    
    public function run(Stream $stream) : Stream
    {
        $stream->setResponse($this->dispatcher($stream->getRequest('uri')));
        return $stream;
    }

    private function dispatcher(string $uri)
    {
        preg_match('/^(\w*?Controller)\/(\w*?Action)$/', $uri, $matches);
        if (3 === count($matches) && class_exists('\\' . $matches[1]) && method_exists('\\' . $matches[1], $matches[2])) {
            $toController = $matches[1];
            $action       = $matches[2];
        } else {
            $toController = 'IndexController';
            $action       = 'indexAction';
        }
        $controller = Container::make('Controller');
        $paramters  = [1];
        return $controller($toController, $action, $paramters);
    }
}

