<?php
namespace BaseStone\Core;

class BaseView extends BaseRequestType
{
    public function __construct()
    {
        $this->request  = Request::getInstance(); 
        $this->response = Response::getInstance();
        $this->type     = 'view';
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

