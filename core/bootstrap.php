<?php

namespace Akf\Core;

abstract class Bootstrap
{
    /**
     *  运行引导 run
     */
    public static function run()
    {
        $a = new Request();

        $stream = Container::make('Stream', $_REQUEST);
        $stream = self::loadComponents(Config::getComponents(), $stream);
        $stream->out();
    }


    private static function loadComponents(array $components, Stream $stream) : Stream
    {
        foreach ($components as $component => $componentConfig) {
            if (class_exists($component)) {
                $config  = Config::getConfig($componentConfig);
                $stream  = Container::make($component, $config)->run($stream);
            }
        }
        return $stream;
    }
}

