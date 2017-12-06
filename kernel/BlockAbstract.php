<?php
namespace Akf\Kernel;

abstract class Block 
{
    protected $name;
    protected $alias;
    protected $func = [];

    abstract protected function init();

    abstract protected function in(Glue $glue);

    abstract protected function out() : Glue;

    public function create()
    {
    
    }
}
