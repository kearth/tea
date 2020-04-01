<?php

namespace Tea\Framework;

/**
 * 环境类
 */
class Env {
    public static function get() {
        return Config::get('env', 'env');
    }
}

