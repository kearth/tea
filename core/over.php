<?php
namespace Akf\Core;

class Over extends Base
{
    public function __construct()
    {
        $this->app = new Container();
        $this->response = Response::getInstance();
    }   

    public static function run()
    {
        $self = self::getInstance();
        $self->response->setOutput('outputviews');
        $self->app->make($self->response->getOutput(), $self->response->getParams());
    }
}

