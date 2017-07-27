<?php
namespace BaseStone\Core;

class BaseSysCMPT
{
    protected $request;
    protected $response;
    protected $sys_cmpt_name;

    public function __construct()
    {
        $this->request  = Request::getInstance();
        $this->response = Response::getInstance();
    }

    public static function create()
    {
        return new static();  
    }
}

