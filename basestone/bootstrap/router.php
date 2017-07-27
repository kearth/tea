<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\BaseSysCMPT;

class Router extends BaseSysCMPT
{
    
    public function route()
    {
        return $this->request;
    }

}

