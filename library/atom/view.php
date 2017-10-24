<?php

namespace Akf\Library\Atom;

use Akf\Core\ReturnValue;

class View extends ReturnValue
{

    public function get() : \Closure
    {
        return function () {
            ob_start();
            echo "<b>Hello world</b><input type='text' />";            
            ob_end_flush();
        };   
    }

    public function set(array $value)
    {
           
    }
}
