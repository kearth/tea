<?php
namespace Akf\Core;

class Container extends Singleton
{
    //存储singleton实例
    private $instance = [];
    //存储bind方法
    private $bind     = [];

    public function bind($abstract, $concrete = null, $shared = false)
    {
        $concrete = $this->getClosure($abstract, $concrete);
        $this->bind[$abstract] = compact("concrete", "shared");
    }

    public function getClosure($abstract, $concrete)
    {
        //if (is_null($concrete)) {
               
        //}

        if ($concrete instanceof \Closure) {
            return $concrete;
        } 

        if (is_object($concrete)) {
            return function ($paramters) {
                return $concrete;
            };
        }

        if (is_string($concrete)) {
            if (class_exists($concrete)) {
                return function ($paramters) {
                    return new $concrete();
                };
            }   
            throw new \Exception("没有这个类");
        }
    }


    /**
     * @name singleton 绑定容器为单例
     * @param $abstract 类名
     * @param $concrete 绑定方法
     * @return void
     */   
    public function singleton($abstract, $concrete)
    {
        $this->bind($abstract, $concrete, true);
    }

    /**
     *  @name instance 绑定实例到容器
     *  @param $abstract 类名
     *  @param $concrete 绑定方法
     *  @return void
     */
    public function instance($abstract, $concrete)
    {
        $this->bind($abstract, $concrete);   
    }

    public function make($abstract, $paramters = [])
    {
        if (array_key_exists($abstract, $this->bind)) {
            if (true == $this->bind[$abstract]['shared']) {
                if (isset($this->instance[$abstract])) {
                    return $this->instance[$abstract];
                }
                $concrete = $this->bind[$abstract]['concrete'];
                $this->instance[$abstract] = $concrete($paramters);
                return $this->instance[$abstract];
            } 
            $concrete = $this->bind[$abstract]['concrete'];
            return $concrete($paramters);
        }
    }
}

