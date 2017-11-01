<?php

namespace Akf\Core\BaseComponent;

use Akf\Core\BaseSource\Stream;

class Back extends BaseComponent
{
    public function run(Stream $stream) : Stream
    {
        $stream->getResponse()->run();
        return $stream;
    }

}
