<?php
namespace BaseStone\Bootstrap;

class Config
{
    private static $instance = null;
    private $allConfig = null;

    private function __construct()
    {
    }
    
    private function __clone()
    {
    }

    private function __sleep()
    {
    }

    private function __wakeup()
    {
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }
    
    public function load($confFile)
    {
        if (file_exists($confFile)) {
            $confInfo = parse_ini_file($confFile,true);
            if ($confInfo) {
                $env = $confInfo['env']['akf_env'];
                $curEnvConf = $confInfo[$env];
                if (isset($curEnvConf['afk_constant'])) {
                    $this->defineconstant($curEnvConf['afk_constant']);
                    unset($curEnvConf['afk_constant']);
                }
                $this->allConfig = $curEnvConf;
            }
        }
    }

    private function defineConstant($constantArr)
    {
        foreach($constantArr as $key => $value) {
            define($key,$value);
        }
    }

    public function get($confName)
    {
        
    }

    public function set($confName,$confValue)
    {
    
    }

}

