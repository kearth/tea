<?php
// 定义初始化常量
define("CONFIG_KEY_INIT", "init");
define("CONFIG_KEY_APP", "app");
define("CONFIG_KEY_CACHE", "cache");
define("CONFIG_KEY_DB", "db");
define("CONFIG_KEY_ROUTE", "route");
define("CONFIG_KEY_AUTOLOAD", "autoload");
define("CONFIG_KEY_ERROR", "error");
define("CONFIG_KEY_LOG", "log");
define("CONFIG_PATH", ROOT_PATH . "/config");

// 定义配置
return [
    //dir or array
    CONFIG_KEY_INIT  => [
        "hello" => "world"
    ],
    CONFIG_KEY_APP   => CONFIG_PATH . "/app.php",
    CONFIG_KEY_CACHE => "",
    CONFIG_KEY_ROUTE => "",
    CONFIG_KEY_ERROR => "",
    CONFIG_KEY_AUTOLOAD => CONFIG_PATH . "/autoload.php"
];
