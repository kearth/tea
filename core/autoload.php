<?php
namespace Akf\Core;

class Autoload
{
    //配置文件加载策略
    const LOAD_CONF_STRATEGY_DEFAULT     = 0;   //默认
    const LOAD_CONF_STRATEGY_SELF_DESIGN = 1;   //自定义
    const LOAD_CONF_STRATEGY_MULTIPLE    = 2;   //多重混合


    private $has_required_file = [];

    public static function register()
    {
        spl_autoload_register(array(new static(),'loadClass'));
    }

    public function loadClass($class)
    {
        $class_file = ROOT_PATH . DIRECTORY_SEPARATOR . substr(str_replace('\\', '/', $class), 4) . '.php';
    
        if (in_array($class_file,$this->has_required_file)) {
            return;
        }

        if (file_exists($class_file)) {
            require $class_file;
            $this->has_required_file[] = $class_file;
        }
    }


    /**
     * @name loadConfStrategy 加载策略
     * @param $loadConfStrategy
     */
    private function loadConfStrategy(int $loadConfStrategy)
    {
        switch ($loadConfStrategy) {
            case self::LOAD_CONF_STRATEGY_DEFAULT :
                return $this->loadConfStrategyDefault();
            case self::LOAD_CONF_STRATEGY_SELF_DESIGN :
                return $this->loadConfStrategySelfDesign();
            case self::LOAD_CONF_STRATEGY_MULTIPLE :
                return $this->loadConfStrategyMutiple();
            default :
                throw new \Exception("没有这个加载策略");
        }   
    }

    /**
     *
     */
    private function loadConfStrategyDefault()
    {
        $classFile = ROOT_PATH . DIRECTORY_SEPARATOR . substr(str_replace('\\', DIRECTORY_SEPARATOR, $class),3) . '.php';
        return $classFile;
    }

    /**
     *
     */
    private function loadConfStrategySelfDesign()
    {
    
    }
    
    /**
     *
     */
    private function loadConfStrategyMutiple()
    {
    
    }
}

