<?php

namespace App;

use Tea\Framework\Plugin\Router as R;

class Router {

    private static string $toNamespace = __namespace__ . "\Action";

    private static string $toSuffix = "Action";

    private static function getAction(array $uri) : string {
        $uriList = explode("/", $uri[1]);
        return implode("\\", array(self::$toNamespace, ...$uriList)) . self::$toSuffix; 
    }

    public static function register() {
        //R::setDefaultRule();
        
        //R::setRule("([/(.*)", Action\IndexAction::class);
       
        R::setRule("/(.*)", function(array $uri){
            $action = self::getAction($uri);
            if (class_exists($action)) {
                $action::execute();
            } else {
                echo "Hello";
            }
        });
        R::run();
    }
}
