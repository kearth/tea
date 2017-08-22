<?php
namespace Akf\Core;

class Config extends Singleton
{
    private $allConfig  = [];
    private $ini        = [];
    
    public function load() : bool
    {
        if (file_exists(CONFIG_DEFAULT)) {
            $confInfo = $this->loadIni(CONFIG_DEFAULT);
            if ($confInfo) {
                $env              = $confInfo['env']['akf_env'];
                $curEnvConf       = $confInfo[$env];
                $this->allConfig  = $curEnvConf;
                $this->ini        = $curEnvConf['akf_ini'];
                defined('ENV') or define('ENV', $env);
                defined('DEBUG') or define('DEBUG', $curEnvConf['akf_debug']);
                defined('ERRORLEVEL') or define('ERRORLEVEL', $curEnvConf['akf_error_level']);
                $this->loadModules();
                $this->loadContant();
                $this->loadInis();
                return true;
            } 
        }
        return false;
    }

    public function getAllConfig() : array
    {
        return $this->allConfig;
    }

    public function getModules() : array
    {
        return $this->allConfig['modules'];
    }

    private function loadModules()
    {
        $modules = $this->loadIni(CONFIG_PATH_ROOT . $this->ini['module']);       
        foreach ($modules[ENV] as $key => $value) {
            $this->allConfig['modules'][$key] = $value;
        }          
    }

    private function loadInis()
    {
        unset($this->ini['module']);
        unset($this->ini['constant']);

        $inis = $this->ini;
        foreach ($inis as $key => $ini) {
            $iniContent = $this->loadIni(CONFIG_PATH_ROOT . $ini);
            $this->allConfig[$key] = $iniContent;
        }
        unset($this->allConfig['akf_ini']);
    }

    private function loadIni($iniFile) : array
    {
        return parse_ini_file($iniFile, true);
    }

    private function loadContant()
    {
        $constants = $this->loadIni(CONFIG_PATH_ROOT . $this->ini['constant']);
        foreach ($constants[ENV] as $key => $value) {
            defined($key) or define($key,$value);
            $this->allConfig['constants'][$key] = $value; 
        }
    }


    public function appload()
    {
        include ROOT_PATH."/config/app.php";
    }

    /**
     * @name get 获取配置
     * @param $name 配置名
     * @return mixed
     */
    public function get($name)
    {
        if (array_key_exists($name, $this->allConfig)) {
            return $this->allConfig[$name];
        }
    }

    /**
     * @name set 设置配置
     * @param $name 配置名
     * @param $value 配置值
     * @return void
     */
    public function set($name,$value)
    {
        $this->allConfig[$name] = $value;
    }
}

