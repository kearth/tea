<?php

namespace Tea\Framework\Flow;

use Tea\Framework\Register as Register;

abstract class Flow {

    protected string $key;

    abstract public function init(array $flow);

    public function getFlow(array $flow) : array {
        if (isset($flow[$this->key])) {
            return $flow[$this->key];
        }  
        throw new Error("no this flow");
    }

    public static function getInstance() : object {
        $instance = Register::getInstance(get_called_class()); 
        if (is_null($instance)) {
            throw new \Error("flow not added!");
        }
        return $instance;
    }

}
