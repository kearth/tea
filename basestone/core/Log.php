<?php
namespace BaseStone\Core;

class Log
{
    private static $instance = null;

    private function __construct()
    {
    
    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    public static function Info($data){
        $log = new self();
        $message = "[".date('Y-m-d H:i:s')."] ".print_r($data,true)."\n";
        return file_put_contents($log->path.$log->file.'.log',$message,FILE_APPEND);
    }

    public function init()
    {
    
    }
}
