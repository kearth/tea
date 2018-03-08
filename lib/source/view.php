<?php

namespace Akf\Library\Source;

use Akf\Core\BaseSource\Response;

class View extends Response
{
    private $viewPath;
    private $param;
    
    public function run()
    {
        ob_start();
        include $this->viewPath;
        ob_end_flush();
    }

    public function defaultRule()
    {
        //$this->rule = function () use ($this){
            //$this->viewPath = ROOT_PATH . '/application/view' . '/';
        //}; 
    }

    public function set(array $value)
    {
        $this->viewPath = ROOT_PATH . '/application/view/index/index.php';
        $this->param = $value;     
    }

    

}
