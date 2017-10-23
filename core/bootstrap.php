<?php
namespace Akf\Core;

abstract class Bootstrap
{
    private static $components = [];

    /**
     *  运行引导 run
     *
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

        $stream = new Stream($_REQUEST);

        foreach (self::$components as $component) {
            $stream = $component->run($stream);
        }
        
        $stream->flowAway();
    }
}

