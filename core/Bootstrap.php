<?php

namespace Tea\Core;

class Bootstrap
{
    public static function run() : void
    {
        Config::init(ROOT_PATH . "/config/app.php");

        Env::init(Config::getConfig("env"));

       // $component = Config::getConfig("component");
       // $classMap = Config::getConfig("classMap");
        Request::run();

        Router::run();
        
        Dispatcher::run();

        Response::run();
    }
}
