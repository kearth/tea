<?php

namespace Tea\Core;

use Tea\Core\Base\Alias;
use Tea\Core\Base\Stream;

class Application
{

    public static function bootstrap()
    {
        echo __class__;
        echo Tea\Core\Base\Stream::class;
        static::seed();
        static::germinate();
        static::grow();
        static::fade();
    }

    private static function seed()
    {
        Stream::from('Mode')->detectWhichMode()->initTheMode("666");
    }

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
