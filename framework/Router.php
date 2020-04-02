<?php

namespace Tea\Framework;

/**
 * 路由类
 */
class Router {

    // 默认路由
    const DEFAULT_ROUTER = 'Index_index';
    // 路由规则映射
    const RULE_MAP = "RuleMap";
    // 路由
    const RULE = "Rule";

    /**
     * 路由map
     */
    public static $routerMap = array();

    /**
     * 路由规则map
     */
    public static $ruleMap = array();

    /**
     * 初始化
     */
    public static function init() : void {
        $conf = Config::initConf(__CLASS__);
        static::$ruleMap = $conf[static::RULE];
        static::$routerMap = $conf[static::RULE_MAP];
    }

    /**
     * 获取路由
     */
    public static function getRouter() : string {
        $server = Request::getServer();
        $uri = $server['REQUEST_URI'];
        // 优先从map中取出对应的分发路由
        if (isset(static::$routerMap[$uri])) {
            return static::$routerMap[$uri];
        }
        // 正则规则匹配路由
        foreach(static::$ruleMap as $rule => $func) {
            $router = preg_replace_callback($rule, $func, $uri);
            if ($router) {
                return $router;
            }
        }
        // 默认路由
        return static::DEFAULT_ROUTER;
    }
}
