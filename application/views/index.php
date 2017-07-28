<?php
namespace Application\Views;

use BaseStone\Core\BaseView;

class Index extends BaseView
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

