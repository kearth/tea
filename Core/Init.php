<?php
namespace core;

/**
 * 类名：Init
 * 功能：初始化整个框架
 * @author:jiakun<kunsama@163.com>
 * @version:1.0
 */
class Init{
    public static $classRegistered = array();   //类注册表
    public static $registerTable = array();     //全局模块注册表
    public static $classRange = array(
        'Controller' => CONTROLLER,
        'Core'       => CORE,
        'Lib'        => LIB,
        'Model'      => MODEL
    );

    /**
     * 函数名：run
     * 功能：运行框架
     * @version:1.0
     * @param void
     * @return void
     */
    public static function run(){
        //$u = new \model\User();
        //var_dump($u);
       // $n = new DBTable('user',
       //     array(
       //         'id'=>'INT UNIQUE',
       //         'userName' =>'char(255)'
       //     ),
       //     array(
       //         'index1' =>'id' ,
       //         'index2' => 'userName',
       //     )
       // );
       // $db = DBConnect::getInstance();
       // $db->createTable($n);
        //路由
        $route = new \core\Route();
        $controllerClass = $route->controller;
        $action = $route->action;
        $request = new \core\Request();
        $response = new \core\Response();
        //控制器
        try{
            $controller = new $controllerClass();
            if(!method_exists($controller,$action)){
                $action = 'index';
            }
            $controller->$action($request,$response);
        } catch(Exception $e){
            echo $e->getMessage();
        }
    }

    /**
     * 函数名：autoLoad
     * 功能：自动加载
     * @version:1.0
     * @param string $class
     * @return boolean
     */
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
