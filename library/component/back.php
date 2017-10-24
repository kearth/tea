<?php

namespace Akf\Library\Component;

use Akf\Core\{Component, Stream, ReturnValue};

class Back extends Component
{
    public function run(Stream $stream) : Stream
    {
        $this->back($stream->getResponse('back'));
        return $stream;
    }

    private function back(ReturnValue $returnValue)
    {
        $closure = $returnValue->get();
        $closure();
    }

}
