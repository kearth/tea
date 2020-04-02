<?php

namespace Tea\Framework;

/**
 * 环境类
 */
class Env {

    /**
     * 获取当前环境
     */
    public static function get() : string {
        return Config::initConf(__CLASS__);
    }
}

