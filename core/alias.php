<?php

namespace Akf\Core;

/**
 *  类别名 Alias
 */
class Alias
{
    /**
     *  加载类别名 load
     *
     *  @param $cfg 加载类别名配置
     *  @return void
     */
    public static function load(array $cfg)
    {
        foreach ($cfg as $name => $alias) {
            class_alias($name, $alias);
        }   
    }
}
