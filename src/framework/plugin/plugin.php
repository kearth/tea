<?php 
namespace Tea\Framework\Plugin;

use Tea\Framework\Register as Register;

class Plugin {

    public function add() {
        $this->isAllow = true;
    }

    public static function getInstance() : object {
        $instance = Register::getInstance(get_called_class()); 
        if (is_null($instance)) {
            throw new \Error("plugin not added!");
        }
        return $instance;
    }

}


