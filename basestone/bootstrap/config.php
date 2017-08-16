<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Singleton;

class Config extends Singleton
{
    private $allConfig = null;

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

    public function appload()
    {
        include ROOT_PATH."/config/app.php";
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

