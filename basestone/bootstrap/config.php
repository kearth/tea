<?php
namespace BaseStone\Bootstrap;

class Config
{
    private static $instance = null;
    private $all_conf = null;

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
    
    public function load($conf_file)
    {
        if (file_exists($conf_file)) {
            $conf_info = parse_ini_file($conf_file,true);
            if ($conf_info) {
                $env = $conf_info['env']['akf_env'];
                $cur_env_conf = $conf_info[$env];
                if (isset($cur_env_conf['afk_constant'])) {
                    $this->defineconstant($cur_env_conf['afk_constant']);
                    unset($cur_env_conf['afk_constant']);
                }
                $this->all_conf = $cur_env_conf;
            }
        }
    }

    private function defineConstant($constant_arr)
    {
        foreach($constant_arr as $key => $value) {
            define($key,$value);
        }
    }

    public function get($conf_name)
    {
        
    }

    public function set($conf_name,$conf_value)
    {
    
    }

}

