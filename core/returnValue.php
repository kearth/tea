<?php

namespace Akf\Core;

abstract class ReturnValue
{
    protected $type;
    protected $content;

    public function __construct(\Closure $cfg)
    {
        $this->type    = get_class($this);
        $this->content = $cfg;
    }

    abstract public function show();
}

