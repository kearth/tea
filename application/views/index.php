<?php
namespace Akf\Application\Views;

use Akf\Core\Application;

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

