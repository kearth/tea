<?php
use Akf\Core\Container;

$container = new Container();
$container->singleton("outputviews", function($paramters){
    $view = ROOT_PATH."/public/".$paramters['view'].".html";
    if (file_exists($view)) {
        ob_start();
        include($view);
        $data = $paramters['data'];
        ob_end_flush();
    }
});

$container->singleton("outputapi", function($paramters){
        $output = [
            'code' => 200,
            'msg'  => "访问成功",
            'data' => $paramters
        ];
        echo json_encode($output);
});
