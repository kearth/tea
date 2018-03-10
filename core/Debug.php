<?php

namespace Tea\Core;

class Debug
{
    public static $env = [
        "dev"     => true,
        "test"    => false,
        "product" => false
    ];

    use TeaTrait\Init;
}
