<?php

namespace Akf\Core;

abstract class Bootstrap
{
    /**
     *  运行引导 run
     */
    public static function run()
    {   
        Container::make('Stream', $_REQUEST)
            ->inject('Router', Config::get('route'))       
            ->inject('Dispatcher')
            ->inject('Back')
            ->out();
    }
}

