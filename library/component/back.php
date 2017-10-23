<?php

namespace Akf\Library\Component;

use Akf\Core\{Component, Stream, ReturnValue};
use Akf\Library\Atom;

class Back extends Component
{
    public function run(Stream $stream) : Stream
    {
        //$this->back($stream->getResponse('back'));
        $a = [
            'res_code' => 1,
            'res_msg'  => 2,
            'result'   => [22332]
        ];
        $this->back(new Atom\Api($a));
        return $stream;
    }

    private function back(ReturnValue $returnValue)
    {
        $returnValue->show();
    }
}
