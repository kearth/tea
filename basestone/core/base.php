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



}

