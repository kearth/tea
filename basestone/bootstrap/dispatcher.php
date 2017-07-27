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
        } else {
            echo "请求不存在";
        }
        $method->output();
    }

}

