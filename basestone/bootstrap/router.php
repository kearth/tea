<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Singleton;

class Router extends Singleton
{
    public $request_source_type;
    public $request_source_name;
    public $request_source_params;
    public $request_method;

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
        $this->request_method = $_SERVER['REQUEST_METHOD'];
        $action  = $_REQUEST['action'];
        $action_arr = explode('/',$action);
        $this->request_source_type = $action_arr[0];
        $this->request_source_name = $action_arr[1];
        $this->request_source_params = $_REQUEST;
        return self::getInstance();
    }


}
