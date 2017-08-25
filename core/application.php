<?php
namespace Akf\Core;

abstract class Application extends Base
{
    protected $request;
    protected $response;
    protected $requestType;

    const REQUEST_TYPE_API       = 'api';
    const REQUEST_TYPE_VIEWS     = 'views';
    const REQUEST_TYPE_RESOURCES = 'resources';


    protected function error(){
        echo "404，老铁，根本就没有这个页面";
    }

    protected function before()
    {
    
    }

    protected function after()
    {
    
    }

    public function __construct()
    {
        parent::__construct();
        $this->request  = Request::getInstance(); 
        $this->response = Response::getInstance();
    }

    public function run()
    {
        $this->before();
        $this->after();
    }

    public function output()
    {
        $params = $this->response->getParams();
        $view = ROOT_PATH."/public/".$params['view'].".html";
        if (file_exists($view)) {
            ob_start();
            include($view);
            $data = $params['data'];
            ob_end_flush();
        }
    }

}

