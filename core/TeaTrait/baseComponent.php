<?php

namespace Akf\Core\BaseComponent;

use Akf\Core\BaseSource\Stream;

abstract class BaseComponent
{
    abstract public function run(Stream $stream) : Stream;

    public function __construct(array $cfg)
    {
    
    }
}
