<?php
namespace Akf\Core;

abstract class Provider
{
    protected $bootLoad = true;
    protected $app;

    abstract public function register();

    protected function __construct()
    {
        $this->app = Container::getInstance();   
    }

    public static function getInstance($paramters = [])
    {
        return new static($paramters);
    }

}
