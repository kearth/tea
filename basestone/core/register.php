<?php
namespace BaseStone\Core;

class Register extends Singleton
{
    private $modules   = [];
    private $instance  = [];
    private $binds     = [];
    private $alias     = [];
    private $event     = [];

    public function bind($abstract, $concrete, $shared)
    {
        if ($concrete instanceof \Closure) {
            $this->binds[$abstract] = $concrete;
        }
        
        if (is_object($concrete)) {
            $this->instance[$abstract] = $concrete;
        }
    }

    public function modules($abstract, $concrete, $orderby)
    {
        
    }

    public function singleton($abstract, $concrete)
    {
        $this->binds($abstract, $concrete, true);
    }

    public function instance($abstract, $concrete)
    {
        $this->binds($abstract, $concrete, false);   
    }

    public function make($abstract, $paramters = [])
    {
        
    }

}

