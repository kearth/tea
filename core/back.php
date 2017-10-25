<?php

namespace Akf\Core;

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
