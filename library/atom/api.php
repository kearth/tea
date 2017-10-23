<?php

namespace Akf\Library\Atom;

use Akf\Core\ReturnValue;

class Api extends ReturnValue
{
    public function show()
    {
        $b = $this->content;       
        $b();
    }
}
