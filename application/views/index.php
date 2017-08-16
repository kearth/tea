<?php
namespace Application\Views;

use BaseStone\Core\Application;

class Index extends Application
{
    public function getAction()
    {
        $params = [
            'view' => 'templates/index',
            'data' => [
                'asfsd',
                'bab'
            ]
        ];
        $this->response->setParams($params);
    }
}

