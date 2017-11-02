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
        $matches = explode('/', $stream->getUri());
        if (2 === count($matches)) {
            $stream->setController($matches[0]);
            $stream->setAction($matches[1]);
        } else {
            $stream->setController('Index');
            $stream->setAction('index');
        }
        $controller = Container::make('Controller');
        return $controller($stream);
    }


}

