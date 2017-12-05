<?php
namespace Akf\Kernel;

abstract class Block 
{
    protected $name;
    protected $alias;
    protected $func = [];

    abstract protected function init();

}
