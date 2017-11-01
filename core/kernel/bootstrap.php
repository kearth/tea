<?php

namespace Akf\Core\Kernel;

use Akf\Core\BaseSource\Stream;

abstract class Bootstrap
{
    /**
     *  运行引导 run
     */
    public static function run()
    {
        $request  = Container::make('Request');
        $stream = Container::singleton('Stream', $request);
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

