<?php
namespace BaseStone\Core;

class Core
{
    protected function toObj($arr)
    {
        $obj = new \StdClass();

        foreach ($arr as $key => $value) {
            $obj->{ strtolower($key) } = $value;    
        }      

        return $obj;
    }
}


