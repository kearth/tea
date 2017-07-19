<?php
namespace Bootstrap;

class Autoload
{
    public function register()
    {
        spl_autoload_register(array($this,'loadClass'));
    }

    public function loadClass(String $class)
    {
        echo $class;
    }
}

