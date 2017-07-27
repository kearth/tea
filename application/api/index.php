<?php
namespace Application\Api;

use BaseStone\Core\BaseApi;

class Index extends BaseApi
{
    public function getAction()
    {
        $this->response->setParams(['abc']);
    }
}
