<?php

namespace Tea\Core;

use Tea\Core\Base\{
    Stream,Alias
};

class Application
{

    public static function bootstrap()
    {
        static::seed();
        static::germinate();
        static::grow();
        static::fade();
    }

    private static function seed()
    {
        Mode::detectWhichMode();
        Alias::set([
            'Mode' => '\Tea\Core\Mode'
        ]);

        Stream::from('Mode')->detectWhichMode()->initTheMode();

    }

    //public static function whichMode()
    //{
        //return php_sapi_name();
    //}

    private static function germinate()
    {
    
    }

    private static function grow()
    {
    }

    private static function fade()
    {
    
    }
}
