<?php

/**
 * 路由配置
 */

// 配置项
retur [
    '{controller|[a-zA-Z]+}/{action}' => function ($controller, $action) {
        return ucfirst($controller) . "/" . $action;
    }
];
