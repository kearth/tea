<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Request;
use BaseStone\Core\Response;

class Router extends Base
{

    private $request;
    private static $route = [];

    public function __construct()
    {
        $this->request  = Request::getInstance();
    }

    public function run()
    {
        $this->routerStartUp();
        $this->router();
        $this->routerShutDown();   
    }

    public function router()
    {

        //self::get('views/ihehehe/get/5', function(){
            //return "hello world";       
        //});
        if (!$this->hasRuledRouter()) {
            $this->defaultRouter();
        }
    }

    public function test(\Closure $function)
    {
        $method  = \Closure::bind($function, $this, get_class());
        var_export($method());
    }


    public function routerStartUp()
    {
    }

    public function routerShutDown()
    {
    
    }
    
    public static function options()
    {
    
    }

    public static function head()
    {
    }

    public static function get(string $name, callable $action)
    {
        if (is_callable($action)) {
            self::$route[$name] = $action();   
        }
    }

    public static function post()
    {
    
    }

    public static function put()
    {
    }

    public static function delete()
    {
    
    }

    public static function trace()
    {
    }

    public static function connect()
    {
    
    }

    public function hasRuledRouter()
    {
        foreach (self::$route as $router => $rule) {
            if (array_key_exists($this->request->getAction(), self::$route)) {
                $this->request->setProvider(self::$route[$this->request->getAction()]);
                return true;
            }
        }
        return false;
    }

    public function defaultRouter()
    {
        $provider = "Application\\".str_replace('/','\\',$this->request->getAction());
        $this->request->setProvider($provider);
    }


}

