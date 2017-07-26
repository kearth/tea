<?php
namespace Application\Api;

use BaseStone\Core\BaseRequestType;

class Index extends BaseRequestType
{
    public function getAction()
    {
        $this->response->setResponse('api',['abc']);
    }
}
