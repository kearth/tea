<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Request;

class Router
{

    public $request_source_type;
    public $request_source_name;
    public $request_source_params;
    public $request_method;
    private static $instance = null;

    private function __construct()
    {
    
    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    //public function __construct(){
        //if(isset($_SERVER['SCRIPT_NAME'])){
            //Log::Info($_SERVER['SCRIPT_NAME']);
            //list($controller,$action) = $this->_getAction($_SERVER['SCRIPT_NAME']);
            //$this->controller = "\controller".$controller."Controller";
            //$this->action = $action;
        //} else {
            //$this->controller = 'Index';
            //$this->action = 'error';
        //}
    //}
    
    private function _getAction($params){
        $controller = '';
        $action = '';
        $result = explode('/',ltrim($params,'/'));
        if(sizeof($result)>1){
            $action = array_pop($result);
            foreach($result as $item){
                $controller .= '\\'.ucfirst($item);
            } 
        } else {
            $controller = '\\Index';
            $action = 'Index';
        }
        return array($controller,$action);
    }
    
    public function getRequest()
    {
        return Request::getInstance()->getRequest();
    }


}
