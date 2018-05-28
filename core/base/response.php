<?php

namespace Akf\Core\BaseSource;

abstract class Response
{
    protected $rule;

    abstract public function run();
    abstract public function defaultRule();

    public function setRule(\Closure $callback)
    {
        $this->rule = $callback;
    }

}

