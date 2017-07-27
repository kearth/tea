<?php
namespace BaseStone\Core;

abstract class BaseRequestType
{
    protected $request;
    protected $response;
    protected $type;

    abstract public function __construct();

    abstract protected function output();

    public function error(){
        echo "404，老铁，根本就没有这个页面";
    }

    public function before()
    {
    
    }

    public function after()
    {
    
    }
    
}
