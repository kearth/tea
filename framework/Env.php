<?php

namespace Tea\Framework;

class Env {
    public static function get() {
        return Config::get('env', 'env');
    }
}

