<?php

namespace Akf\Core;

abstract class BaseComponent
{
    abstract public function run(Stream $stream) : Stream;

    public function __construct(array $cfg)
    {
    
    }
}
