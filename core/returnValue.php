<?php

namespace Akf\Core;

abstract class ReturnValue
{
    protected $type;
    protected $content;

    public function __construct()
    {
        $this->type = get_class($this);
    }

    abstract public function get() : \Closure;
    abstract public function set(array $value);
}

