<?php

namespace Tea\Framework;

class Router {

    public static $routerMap = array();

    public static $ruleMap = array();

    public static function init() {
        static::$ruleMap = Config::get('rule', 'router');
        static::$routerMap = Config::get('map', 'router');
    }

    public static function getRouter() {
        $server = Request::getServer();
        $uri = $server['REQUEST_URI'];
        if (isset(static::$routerMap[$uri])) {
            return static::$routerMap[$uri];
        }
        foreach(static::$ruleMap as $rule => $func) {
            $router = preg_replace_callback($rule, $func, $uri);
            if ($router) {
                return $router;
            }
        }
        return 'Index';
    }
}
