<?php
//全局常量
define('HTTP_OK', 200);
//HTTP_OK = 200
//HTTP_MOVED = 301
//HTTP_FOUND = 302
//HTTP_NOT_MODIFIED = 304
//HTTP_BAD_REQUEST = 400
//HTTP_FORBIDDEN = 403
//HTTP_NOT_FOUND = 404
//HTTP_METHOD_NOT_ALLOWED = 405
//HTTP_INTERNAL_SERVER_ERROR = 500
//HTTP_NOT_IMPLEMENTED = 501
//HTTP_BAD_GATEWAY = 502
//HTTP_GATEWAY_TIMEOUT = 504
//HTTP_VERSION_NOT_SUPPORTED = 505

//系统配置
return [
    'env'  => 'dev',
    'bind' => CONFIG_PATH_ROOT . 'bind.php', 
    'alias' => CONFIG_PATH_ROOT . 'alias.php',
    'components' => [
        'Akf\Library\Component\Router'     => [
            'level' => 1,
            'cfg'   => CONFIG_PATH_ROOT . 'route.php'
        ],
        'Akf\Library\Component\Dispatcher' => [
            'level' => 2
        ],
        'Akf\Library\Component\Back'       => [
            'level' => 3
        ]

    ],

];
