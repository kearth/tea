<?php

namespace Tea\Core;

class Env
{

    use TeaTrait\Init;

    public static function get(string $envName) : string
    {
        return getenv($envName);   
    }

    public static function set(string $envName, string $value)
    {
        putenv($envName . "=" . $value);
    }
}

