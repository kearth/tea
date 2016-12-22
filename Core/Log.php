<?php
namespace core;

class Log{
    public $path;
    public $file;
    public function __construct(){
        $conf = \core\Config::getConfig('Log');
        $this->path = $conf['path'];
        $this->file = $conf['file'];
        if(!is_dir($this->path)){
            mkdir($this->path,'0777',true);
        }
    }

    public static function  Info(){
        $log = new self();
        $message = date('Y-m-d H:i:s');
        return file_put_contents($log->path.$log->file.'.log',json_encode($message),FILE_APPEND);
    }
}
