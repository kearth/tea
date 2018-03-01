<?php

namespace Tea\Kernel;

class Env
{
    private static $envList = [];

    public function __construct(array $envList)
    {
        self::$envList = $envList;
    }

    public staitc function get(string $envName)
    {
        if (isset(self::$envList[$envName])) {
            return self::$envList[$envName];
        }
    }
}
