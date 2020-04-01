<?php

/**
 * 路由配置
 */

// 配置项
return array(
    'rule' => array(
        '/{controller|[a-zA-Z]+}\/{action}/' => function ($matches) {
            $controller = $matchs[0];
            $action = $matchs[1];
            return ucfirst($controller) . "/" . $action;
        },
    ),
    'map' => array(
        '/service/index' => 'controller/index'
    ),
);
