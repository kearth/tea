<?php

namespace Tea\Core;

class Seed
{

    public static function sow()
    {
        return $this;   
    }

    public function germinate()
    {
        call_user_func_array($this->closure, $this->args);
    }


}
