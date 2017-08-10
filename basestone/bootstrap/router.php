<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Request;
use BaseStone\Core\Response;

class Router extends Base
{

    private $request;
    private $response;

    public function __construct()
    {
        $this->request  = Request::getInstance();
        $this->response = Response::getInstance();
    }

    public function run()
    {
        $this->routerStartUp();
        $this->router();
        $this->routerShutDown();   
    }

    public function router()
    {
        return $this->request;
    }

    public function routerStartUp()
    {

    }

    public function routerShutDown()
    {
    
    }

}

