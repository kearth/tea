<?php

namespace Tea\Core;

class Application
{
    public static function bootstrap()
    {
        Seed::sow();
    }

    public static function whichMode()
    {
        return php_sapi_name();
    }
}
