<?php
namespace BaseStone\Bootstrap;

use BaseStone\Core\Base;
use BaseStone\Core\Response;

class Over extends Base
{
    
    public function run()
    {
        $this->response = Response::getInstance();
        $bac = $this->response->getParams();
    }

}
