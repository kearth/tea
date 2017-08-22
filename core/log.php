<?php
namespace BaseStone\Core;

class Log extends Singleton
{
    public static function Info($data){
        $log = new self();
        $message = "[".date('Y-m-d H:i:s')."] ".print_r($data,true)."\n";
        return file_put_contents($log->path.$log->file.'.log',$message,FILE_APPEND);
    }

}
