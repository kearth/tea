<?php
namespace core;

class Init{
    public static $classRegistered = array();
    public static $classRange = array(
        'Controller' => CONTROLLER,
        'Core'       => CORE,
        'Lib'        => LIB,
        'Model'      => MODEL
    );

    public static function run(){
        $route = new \core\Route();
        $controllerClass = $route->controller;
        $action = $route->action;
        $request = new \core\Request();
        $response = new \core\Response();
        try{
            $controller = new $controllerClass($request,$response);
            $controller->$action();
        } catch(Exception $e){
            echo $e->getMessage();
        }
    }

    public static function autoLoad($class){
        $class = basename(str_replace('\\','/',$class));
        $classExist = false;
        if(array_key_exists($class,self::$classRegistered)){
            $classExist = true;
        } else {
            foreach(self::$classRange as $dir){
                $file = $dir.'/'.$class.'.php';
                if(is_file($file)){
                    include_once($file);
                    self::$classRegistered[$class] = $class;
                    $classExist = true;
                }
            }
        }
        return $classExist;
    }
}
