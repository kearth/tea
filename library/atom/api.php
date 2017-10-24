<?php

namespace Akf\Library\Atom;

use Akf\Core\ReturnValue;

class Api extends ReturnValue
{
    public function set(array $value)
    {
        if (isset($value['code']) && isset($value['msg']) && isset($value['data'])) {
            $this->content = $value;
        } else {
            throw new \Exception('paramters is wrong');  
        }
    }

    public function get() : \Closure
    {
        $code = $this->content['code'];
        $msg  = $this->content['msg'];
        $data = $this->content['data'];
        return function () use ($code, $msg, $data) {
            $output = [
                'code' => $code,
                'msg'  => $msg,
                'data' => $data
            ];
            echo json_encode($output);
        };
    }
}
