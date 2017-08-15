<?php
namespace BaseStone\Core;

abstract class Base 
{
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

    protected $container; 


    public function __invoke(string $name)
    {
        $this->test($name);
    }

    //public static function __set_state(array $properties)
    //{
        //return error_log(print_r($properties,1));
    //}

    //protected function init()
    //{
        //foreach ($this->init as $method) {
            //call_user_func($method_exists)
        //}
    //}

}


