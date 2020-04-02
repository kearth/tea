<?php

/**
 * 路由配置
 */
return array(
    Tea\Framework\Router::RULE => array(
        '|\/(.*)\/(.*)\?.*|' => function ($matches) {
            $controller = $matches[1];
            $action = $matches[2];
            return $controller . "_" . $action;
        },
    ),
    Tea\Framework\Router::RULE_MAP => array(
        '/service/index' => 'controller/index_action'
    ),
);
