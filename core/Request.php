<?php

namespace Tea\Core;

use Tea\Core\TeaInterface\LoggerInterface;

class Request
{
    public static function run()
    {
    
    }

    public function __construct(LoggerInterface $logger)
    {
        var_export($logger);
    }

}
