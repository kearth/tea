<?php
namespace Akf\Library;

use Akf\Core\Component;
use Akf\Core\Stream;

class Router extends Component
{
    public function __construct(\closure $cfg)
    {
        var_dump($cfg);
    }


    public function run(Stream $stream) : Stream
    {

        return $stream;
    }

}

