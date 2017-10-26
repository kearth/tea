<?php

namespace Akf\Library\Source;


class View extends \Response
{
    private $viewPath;
    private $param;

    


    public function get() : \Closure
    {
        return function () {
            ob_start();
            include $this->viewPath;
            ob_end_flush();
        };   
    }

    public function set(array $value)
    {
        $this->viewPath = ROOT_PATH . '/application/view/index/index.php';
        $this->param = $value;     
    }

    public function path()
    {
    
    }
}
