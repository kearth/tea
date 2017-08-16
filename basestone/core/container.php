<?php
namespace BaseStone\Core;

class Container
{
    private $register;

    public function __construct()
    {
        $this->register = Register::getInstance();    
    }

    public function __call($name, $paramters)
    {
        return call_user_func_array([$this->register,$name], $paramters); 
    }
       
}

