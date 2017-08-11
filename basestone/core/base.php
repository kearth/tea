<?php
namespace BaseStone\Core;

abstract class Base 
{
 
    public static function create()
    {
        return new static();  
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

    //public static function __set_state(array $properties)
    //{
        //return error_log(print_r($properties,1));
    //}

    public function test(string $test)
    {
        echo $test;
    }

}

