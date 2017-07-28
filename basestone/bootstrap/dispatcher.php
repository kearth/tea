<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\BaseSysCMPT;

class Dispatcher extends BaseSysCMPT
{
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

}

