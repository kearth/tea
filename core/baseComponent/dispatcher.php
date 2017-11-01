<?php

namespace Akf\Core\BaseComponent;

use Akf\Core\Kernel\Container;
use Akf\Core\BaseSource\Stream;

class Dispatcher extends BaseComponent
{
    
    public function run(Stream $stream) : Stream
    {
        $stream->setResponse($this->dispatcher($stream));
        return $stream;
    }

    private function dispatcher(Stream $stream)
    {
        //preg_match('/^(\w*?Controller)\/(\w*?Actio\n)$/', $stream->getUri(), $matches);
        $matches = explode('/', $stream->getUri());
        if (2 === count($matches)) {
            $stream->setController($matches[0]);
            $stream->setAction($matches[1]);
        } else {
            $stream->setController('Index');
            $stream->setAction('index');
        }
        //if (3 === count($matches) && class_exists('\\' . $matches[1]) && method_exists('\\' . $matches[1], $matches[2])) {
            //$toController = $matches[1];
            //$action       = $matches[2];
        //} else {
            //$toController = 'IndexController';
            //$action       = 'indexAction';
        //}
        $controller = Container::make('Controller');
        $paramters  = [1];
        return $controller($toController,);
    }
}

