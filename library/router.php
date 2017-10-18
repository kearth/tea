<?php
namespace Akf\Library;

use Akf\Core\Component;
use Akf\Core\Stream;

class Router extends Component
{
    private static $routeRule = [];

    public function __construct(array $cfg)
    {
        $this->setRouteRule($cfg);
    }


    private function setRouteRule(array $routeRules)
    {
        foreach ($routeRules as $routeRule) {
            if (is_callable($routeRule)) {
                self::$routeRule[] = $routeRule;
            }
        }   
    }

    private function getRouteRule()
    {
           
    }

    private function route(string $uri) : string
    {
           
    }


    public function run(Stream $stream) : Stream
    {
        $uri = 6666;

        return $stream;
    }

}

