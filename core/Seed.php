<?php

namespace Tea\Core;

class Seed
{
    private $closure;
    private $args = [];

    public function __construct(\Closure $closure)
    {
        $args = func_get_args();
        $num  = func_num_args();

        if ($num > 1) {
            array_shift($args);
            $this->args = $args;
        }

        $this->closure = $closure;
    }

    public function germinate()
    {
        call_user_func_array($this->closure, $this->args);
    }


}
