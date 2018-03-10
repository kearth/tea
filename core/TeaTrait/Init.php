<?php

namespace Tea\Core\TeaTrait;

use Tea\Core\Seed;

trait Init
{
    public static function init(Seed $seed)
    {
        $seed->germinate();
    }
}

