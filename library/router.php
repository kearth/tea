<?php
namespace Akf\Library;

use Akf\Core\Component;
use Akf\Core\Stream;

/**
 *
 */
class Router extends Component
{
    private static $routeRule = [];

    public function __construct(array $cfg)
    {
        $this->setRouteRule($cfg);
    }


    private function setRouteRule(array $routeRules)
    {
        foreach ($routeRules as $rule => $routeRule) {
            if (is_callable($routeRule)) {
                self::$routeRule[$rule] = $routeRule;
            }
        }   
    }

    private function getRouteRule()
    {
        return self::$routeRule;
    }

    private function route(string $uri) : string
    {
        $routeRules = $this->getRouteRule();
        foreach ($routeRules as $rule => $routeRule) {
            $rule = $this->eregWash($rule);
            preg_match_all($rule, $uri, $matches, PREG_PATTERN_ORDER);
            if (1 === count($matches)) {
                return $routeRule();
            } 
           
            if (1 <  count($matches)){
                array_shift($matches);
                return call_user_func_array($routeRule, $matches);
            }
        }
        return $uri;
    }


    public function run(Stream $stream) : Stream
    {
        $stream->setRequest('uri', $this->route($stream->getRequest('uri')));
        echo $stream->getRequest('uri');
        return $stream;
    }

    private function eregWash(string $rule) : string
    {
        //$ruleArr = explode('|', $rule);
        $b = [];
        echo $rule;
        echo '<br>';
        $a = preg_replace_callback(
            '/(\{(.*?)\})/i', 
            function ($matches) {
                $ruleArr = explode('|', $matches[2]);
                if (2 ===  count($ruleArr)) {
                    $replacement = $ruleArr[0];
                    $rule        = $ruleArr[1];
                } else {
                    $replacement = $matches[2];
                }
                return "({$replacement})";
            }, 
            $rule
        );
        echo $a;
        exit;


        $ruleReplace = preg_replace('/(\{(.*)\})/i', 
            $replacement, 
            $rule
        );
        return  "/" . addcslashes($ruleReplace, '/') . "/";
    }

}

