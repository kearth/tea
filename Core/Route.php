<?php
namespace core;

class Route{

    public $controller;
    public $action;

    public function __construct(){
        if(isset($_SERVER['SCRIPT_NAME'])){
            list($controller,$action) = $this->_getAction($_SERVER['SCRIPT_NAME']);
            $this->controller = "\controller".$controller."Controller";
            $this->action = $action;
        } else {
            $this->controller = 'Index';
            $this->action = 'Index';
        }
    }
    
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

}
