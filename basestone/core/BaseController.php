<?php
namespace core;

class BaseController{

    protected $request;
    protected $response;

    public function __construct(){
    }

    protected function backToBrowser($response){
        $view = new \core\View($response->getAttribute('view'),$response->getAttribute('data'));
        $view->show();
    }

    public function error(){
        echo "404，老铁，根本就没有这个页面";
    }

    
}
