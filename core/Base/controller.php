<?php
namespace Akf\Core\BaseSource;

use Akf\Core\Kernel\Container;

class Controller
{
    public function __invoke(Stream $stream)
    {
        $controller = $stream->getController() . 'Controller';
        $action     = $stream->getAction() . 'Action';
        $ctr = Container::make($controller);
        $ctr->request   = $stream->getParam();
        $ctr->$action();
        if (isset($ctr->response['type'])) {
            $returnValue = Container::make($ctr->response['type']);
            $returnValue->set($ctr->response);
            return $returnValue;
        }
        throw new \BaseException('response is invailed');
    }
}

