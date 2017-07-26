<?php
namespace BaseStone\Core;

class BaseRequestType
{
    protected $request;
    protected $response;
    protected $type;
    protected static $type_array = [
        'api',
        'resources',
        'views'
    ];

    public function __construct()
    {
        $this->request  = Request::getInstance(); 
        $this->response = Response::getInstance();
    }

    public function before()
    {
    
    }

    public function after()
    {
    
    }


    protected function backToBrowser($response){
        $view = new \core\View($response->getAttribute('view'),$response->getAttribute('data'));
        $view->show();
    }

    public function error(){
        echo "404，老铁，根本就没有这个页面";
    }

    
}
