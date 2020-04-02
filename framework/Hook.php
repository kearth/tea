<?php 

namespace Tea\Framework;

/**
 * 钩子类
 */
class Hook {

    /**
     * 分发前
     */
    public static function before() {
        $func = Config::get('HookBefore');
        $func();
    }

    /**
     * 分发后
     */
    public static function after() {
        $func = Config::get('HookAfter');
        $func();
    }

}
