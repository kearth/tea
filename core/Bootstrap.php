<?php

namespace Tea\Core;

class Bootstrap
{
    public static function run()
    {
        Request::run();

        Router::run();
        
        Dispatcher::run();

        Response::run();
    }
}
