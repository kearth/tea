<?php

namespace Tea\Core;

class Application
{
    public static function bootstrap()
    {
        //判断模式
        static::appMode()->start();
    }

    public static function appMode()
    {
        return new HttpMode();
    }
}
