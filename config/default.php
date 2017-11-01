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
    'env'    => 'test',
    'bind'   => [
        'Api' => function () {
            return new Api();
        },
        'Router' => function ($config) {
            return new Router($config);
        },
    ],
    'bindSingle' => [
        'Stream' => function ($request) {
            return new Stream($request);
        }
    ],
    'alias'  => [
        'Akf\Core\Kernel\Container'   => 'Container',
        'Akf\Core\BaseSource\Stream'      => 'Stream',
        'Akf\Core\Model'       => 'Model',
        'Akf\Core\BaseSource\Response' => 'Response',
        'Akf\Core\BaseSource\Request'  => 'Request',
        'Akf\Core\BaseException' => 'BaseException',
        'Akf\Core\Kernel\Bootstrap' => 'Bootstrap',
        'Akf\Core\BaseSource\Controller' => 'Controller',
        'Akf\Core\BaseComponent\Router' => 'Router',
        'Akf\Core\BaseComponent\Dispatcher' => 'Dispatcher',
        'Akf\Core\BaseComponent\Back' => 'Back',
        'Akf\Library\Source\Api' => 'Api',
        'Akf\Library\Source\View'=> 'View',
    ],
    'components' => [
        'Router'     => CONFIG_PATH_ROOT . 'route.php',
        'Dispatcher' => '',
        'Back'       => ''
    ],

];
