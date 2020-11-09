<?php 

namespace  Tea\App\Action;

abstract class BaseAction {

    abstract public static function execute();

    public static function getParams() {
        return $_REQUEST;
    }

    public static function response(string $output) {
        ob_start(); 
        echo $output;
        ob_end_flush();
    }    

    public static function apiFormat(array $data = array(), int $errno = 0, string $errmsg = "success") : string {
        return json_encode(array(
            "errno"  => $errno,
            "errmsg" => $errmsg,
            "data"   => $data,
        ));
    }

    public static function pageFormat() : string {
        //TODO 
    }
}

