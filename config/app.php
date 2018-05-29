<?php

return [
    "core" => [
        "init" => [
        
        ]
    ]
];


/**
use Tea\Core\{
    Env,Seed,Debug,Config,Container
};

Env::init(new Seed(function($path){
    $dotenv = new \Dotenv\Dotenv($path);
    $dotenv->load();
}, ROOT_PATH));


Debug::init(new Seed(function($debug){
    if (isset(Debug::$env[$debug]) && true == Debug::$env[$debug]) {
        $whoops = new \Whoops\Run;
        $whoops->pushHandler(new \Whoops\Handler\PrettyPageHandler);
        $whoops->register();
    }
}, Env::get("ENV")));

Config::init(new Seed(function($configFileList){
    foreach ($configFileList as $configFile) {
            Config::load($configFile);           
        }
    },
    [
        ROOT_PATH . "/config/route.php"
    ]
));




Container::bind("\Tea\Core\TeaInterface\LoggerInterface", "\Tea\Core\Log");

new Tea\Core\Request(Container::get("\Tea\Core\TeaInterface\LoggerInterface"));

**/

