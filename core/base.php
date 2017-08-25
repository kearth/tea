<?php
namespace Akf\Core;

abstract class Base 
{
    protected $app;
    protected $container; 

    public function __construct()
    {
        $this->app = Container::getInstance();
    }

    public static function getInstance($paramters = [])
    {
        return new static($paramters);
    }

    public function __call($name, $arg)
    {
        error_log(print_r($name,1));
    }

    public function __toString()
    {   
        
        return get_class($this);
    }



    public function __invoke(string $name)
    {
        $this->test($name);
    }

}


