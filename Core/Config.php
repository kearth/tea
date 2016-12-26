<?php
namespace core;

class Config{
    public static $conf = array();

    public static function getConfig($name,$file=ROOT."/conf.php"){
        if(is_file($file)){
            $conf = include $file;
            if(isset($conf[$name])){
                self::$conf[$file] = $conf;
                return $conf[$name];
            } else {
               throw new \Exception('没有这个配置项'.$name);
            }
        } else {
            throw new \Exception('找不到配置文件'.$file);
        }

    }

}
