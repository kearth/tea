<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Response;
use BaseStone\Core\Container;

class Over extends Base
{
    public function __construct()
    {
        $this->app = new Container();
        $this->response = Response::getInstance();
    }   

    public function run()
    {
        $this->response->setOutput('outputviews');
        $this->app->make($this->response->getOutput(), $this->response->getParams());
    }
}

