<?php

namespace Tea\Framework\Plugin;

class GlobalVal extends Plugin {

    private array $global = array();

    public static function get(string $name, $default = null) {
        return self::getInstance()->global[$name] ?? $default;
    }

    public static function set(string $name, $value, bool $force = false) : bool {
        if (isset(self::getInstance()->global[$name]) && !$force) {
            return false;
        }
        self::getInstance()->global[$name] = $value;
        return true;
    }

}

