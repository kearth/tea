<?php
namespace Akf\Core;

class Register extends Singleton
{
    private $singleton  = [];
    private $instance   = [];

    protected function binds($abstract, $concrete, $shared = false)
    {
        if ($shared) {
            if ($concrete instanceof \Closure) {
                $this->singleton[$abstract] = $concrete;
            } else {
                $this->singleton[$abstract] = function($paramters){
                    return $concrete;
                };
            }
        } else {
            if ($concrete instanceof \Closure) {
                $this->instance[$abstract] = $concrete;
            } else {
                $this->instance[$abstract] = function($paramters){
                    return $concrete;
                };
            }
        }
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
        if (array_key_exists($abstract, $this->singleton)) {
            return $this->singleton[$abstract]($paramters);       
        } elseif (array_key_exists($abstract, $this->instance)) {
            return $this->instance[$abstract]($paramters);
        } else {
            throw new \Exception();
        }
    }
}

