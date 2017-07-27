<?php
namespace BaseStone\Core;

class BaseApi extends BaseRequestType
{

    private $output;

    public function __construct()
    {
        $this->request  = Request::getInstance(); 
        $this->response = Response::getInstance();
        $this->type     = 'api';
    }

    public function output()
    {
        $this->output = [
            'code' => 200,
            'msg'  => "访问成功",
            'data' => $this->response->getParams()
        ];
        echo json_encode($this->output);
    }

}
