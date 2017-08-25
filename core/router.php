<?php
namespace Akf\Core;

class Router extends Provider
{

    public function register()
    {
        
    }

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
        if (!$this->hasRuledRouter()) {
            $this->defaultRouter();
        }
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
        $provider = "Akf\\Application\\".str_replace('/','\\',$this->request->getAction());
        $this->request->setProvider($provider);
    }


}

