<?php

namespace Akf\Core;

abstract class BaseSource
{
    public function inject($class, $paramters = [])
    {
        return Container::make($class, $paramters)->run($this);
    }
}
