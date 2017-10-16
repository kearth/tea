<?php
namespace Akf\Library;

use Akf\Core\Component;
use Akf\Core\Stream;

class Dispatcher extends Component
{
    
    public function run(Stream $stream) : Stream
    {
        return $stream;
    }

}

