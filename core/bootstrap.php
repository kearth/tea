<?php

namespace Akf\Core;

abstract class Bootstrap
{
    /**
     *  组件
     */
    private static $components = [];

    /**
     *  运行引导 run
     *
     *  @return void
     */
    public static function run()
    {
        $aliasCfg = Config::getAliasCfg();
        Alias::load($aliasCfg);

        Config::getBindCfg();
        $stream = Container::make('Stream', $_REQUEST);
        $router = Container::make('Router', Config::getComponentCfg(CONFIG_PATH_ROOT . 'route.php'));       
        $stream = $router->run($stream);
        $dispatcher = Container::make('Dispatcher');
        $stream = $dispatcher->run($stream);
        $back = Container::make('Back');
        $stream = $back->run($stream);
        $stream->flowAway();
    }
}

