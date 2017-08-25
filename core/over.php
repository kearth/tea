<?php
namespace Akf\Core;

class Over extends Base
{
    public function __construct()
    {
        parent::__construct();
        $this->response = Response::getInstance();
    }

    public function run()
    {
        $this->response->setOutput('outputviews');
        $this->app->make($this->response->getOutput(), $this->response->getParams());
    }
}

