<?php
namespace Akf\Core;

abstract class Component
{
    abstract public function run(Stream $stream) : Stream;

    public function __construct(array $cfg)
    {
    
    }
}
