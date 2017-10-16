<?php
namespace Akf\Core;

abstract class Bootstrap
{
    private static $components = [];

    /**
     *  @name run 运行
     *  @access public
     *  @return void
     */
    public static function run()
    {
        $components = Config::getComponents();       

        foreach ($components as $componentClass => $component) {
            if (class_exists($componentClass)) {
                $cfg = Config::getComponentCfg($component['cfg'] ?? '');
                self::$components[$component['level']] =  new $componentClass($cfg);
            }
        }

        $stream = new Stream();

        foreach (self::$components as $component) {
            $stream = $component->run($stream);
        }

        echo "hello world";
        //$modules = Config::getInstance()->getModules();

        //foreach ($modules as $module) {
            //if (class_exists($module)) {
                //$module::getInstance()->run();
            //} 
        //}
    }

}

