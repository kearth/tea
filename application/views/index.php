<?php
namespace Application\Views;

use BaseStone\Core\BaseViews;

class Index extends BaseViews
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

