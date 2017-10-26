<?php

namespace Akf\Core;

class Back extends BaseComponent
{
    public function run(Stream $stream) : Stream
    {
        $closure = $stream->getResponse()->get();
        $closure();
        return $stream;
    }

}
