<?php

namespace Akf\Library;

use Akf\Core\Component;
use Akf\Core\Stream;

/**
 *  路由组件 Router
 */
class Router extends Component
{
    private static $routeRule = [];

    /**
     *  构造实体 __construct
     *
     *  @param  array $cfg
     *  @return void
     */
    public function __construct(array $cfg)
    {
        $this->setRouteRule($cfg);
    }

    /**
     *  写入路由规则 setRouteRule
     *
     *  @param  array $routeRules
     *  @return void
     */
    private function setRouteRule(array $routeRules)
    {
        foreach ($routeRules as $rule => $routeRule) {
            if (is_callable($routeRule)) {
                self::$routeRule[$rule] = $routeRule;
            }
        }   
    }

    /**
     *  获取路由规则 getRouteRule
     *
     *  @return array
     */
    private function getRouteRule() : array
    {
        return self::$routeRule;
    }

    /**
     *  路由 route
     *
     *  @param  string $uri
     *  @return string
     */
    private function route(string $uri) : string
    {
        $routeRules = $this->getRouteRule();
        foreach ($routeRules as $rule => $routeRule) {
            $rule = $this->eregWash($rule);
            preg_match_all($rule, $uri, $matches, PREG_PATTERN_ORDER);
            $matches = array_filter(
                $matches,
                function ($var) {
                    return !empty($var) ?? false;     
                }
            );
            if (!empty($matches)) {
                array_shift($matches);
                $matches = array_map(
                    function ($var) {
                        return $var[0];
                    },
                    $matches
                );
                return call_user_func_array($routeRule, $matches);
            }
        }
        return $uri;
    }

    /**
     *  组件运行 run
     *
     *  @param  Akf\Core\Stream $stream
     *  @return Akf\Core\Stream
     */
    public function run(Stream $stream) : Stream
    {
        $stream->setRequest('uri', $this->route($stream->getRequest('uri')));
        return $stream;
    }

    /**
     *  清洗正则表达式 eregWash
     *
     *  @param  string $rule
     *  @return string
     */
    private function eregWash(string $rule) : string
    {
        $ruleReplace = preg_replace_callback(
            '/(\{(.*?)\})/i', 
            function ($matches) {
                $ruleArr = explode('|', $matches[2]);
                if (2 ===  count($ruleArr)) {
                    $ruleMatch   = $ruleArr[1];
                } else {
                    $ruleMatch   = "\w*";
                }
                return "({$ruleMatch})";
            }, 
            $rule
        );
        return  "/^" . addcslashes($ruleReplace, '/') . "$/";
    }
}

