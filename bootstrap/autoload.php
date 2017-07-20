<?php
namespace Bootstrap;

class Autoload
{
    public function register()
    {
        spl_autoload_register(array($this,'loadClass'));
    }

    public function loadClass($class)
    {
        error_log(print_r($class,1));
    }
}

