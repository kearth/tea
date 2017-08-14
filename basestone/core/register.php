<?php
namespace BaseStone\Core;

class Register extends Singleton
{
    private $modules   = [];
    private $singleton = [];
    private $instance  = [];
    private $interface = [];
    private $binds     = [];

    public function bind($abstract, $concrete, $shared)
    {
        if ($concrete instanceof \Closure) {
            $this->binds[$abstract] = $concrete;
        } else {
            $this->instance[$abstract] = $concrete;
        }
    }

    public function singleton()
    {
    
    }

    public function instance()
    {
    
    }

    public function make($abstract, $paramters = [])
    {
        
    }

}

