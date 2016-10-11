<?php
error_reporting(E_ALL);
function autoLoad($class){
    $file = __DIR__."/Controller/".$class.'.php';
    if(is_file($file)){
        require_once($file);
    }
}

spl_autoload_register('autoLoad');

$uri = ltrim($_SERVER['REQUEST_URI'],'/');
$inputData = explode('?',$uri);
if(sizeof($inputData)>1){
    $params = getAllParams($inputData[1]);
}
else{
    $params = array();
}

if(!empty($params)){
    $controller = $params['Controller']."Controller";
    $action = $params['Action'];
    try{
        $a = new $controller();
        $a->$action();
    }
    catch(Exception $e){
        echo $e->getMessage();
    }
}


function getAllParams($paramsRaw){
    $params = explode('&',$paramsRaw);
    $paramsArray = array();
    foreach($params as $value){
        $everParam = explode('=',$value);
        $paramsArray[ucfirst($everParam[0])] = ucfirst($everParam[1]);
    }
    return $paramsArray;
}









