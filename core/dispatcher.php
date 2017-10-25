<?php

namespace Akf\Core;

class Dispatcher extends Component
{
    
    public function run(Stream $stream) : Stream
    {
        $stream->setResponse('back', $this->dispatcher($stream->getRequest('uri')));
        return $stream;
    }

    private function dispatcher(string $uri)
    {
        preg_match('/^(\w*?Controller)\/(\w*?Action)$/', $uri, $matches);
        if (3 === count($matches) && class_exists('\\' . $matches[1]) && method_exists('\\' . $matches[1], $matches[2])) {
            $controller = $matches[1];
            $action     = $matches[2];
        } else {
            $controller = 'IndexController';
            $action     = 'indexAction';
        }
        $ctr = new Controller;
        $paramters = [1];
        return $ctr($controller, $action, $paramters);
    }
}

