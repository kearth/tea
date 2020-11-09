<?php

namespace Tea\Framework\Plugin;

class Router extends Plugin {

    private const DEFAULT_METHOD = "execute";

    private const DEFAULT_SUCCESS = "success";

    private array $ruleList = array();

    public static function setDefaultRule() {
        self::getInstance()->ruleList["/(.*)"] = function(array $uri){
            echo self::DEFAULT_SUCCESS;
        }; 
    }

    public static function setRule(string $uri, $rule) : void {
        if ($rule instanceof \Closure) {
            self::getInstance()->ruleList[$uri] = $rule;  
        } elseif (class_exists($rule)){
            self::getInstance()->ruleList[$uri] = $rule . "::" . self::DEFAULT_METHOD; 
        } else {
            self::getInstance()->ruleList[$uri] = function() use ($rule){
                if (is_string($rule) || is_numeric($rule)) {
                    echo $rule; 
                }
            };
        }
    }

    public static function run() {
        $uri = $_SERVER['REQUEST_URI'];
        $path = parse_url($uri)["path"];
        if(empty(self::getInstance()->ruleList)) {
            self::setDefaultRule();
        }
        $map = self::getInstance()->ruleList;
        foreach($map as $rule => $func) {
            $router = preg_replace_callback('|' . $rule . '|', $func, $path);
            if ($path !== $router) {
                continue;
            }
        }
    }
}

