<?php

namespace Tea\Kernel;

/**
 * PSR-4
 */

class Autoload
{

    protected static $prefixes = [];

    public static function addNamespace(string $prefix, string $baseDir) : void
    {
        if (isset(self::$prefixes[$prefix]) === false) {
            self::$prefixes[$prefix] = $baseDir;
        }
    }

    public static function register()
    {
        spl_autoload_register([__CLASS__, 'loadClass']);
    }

    protected static function loadMappedFile($prefix, $class)
    {

    }

    protected static function requireFile(string $file) : bool
    {
        if (is_file($file)) {
            require($file);
            return true;
        }
        return false;
    }

    protected static function loadClass(string $class) : void
    {
        if (false !== $pos = strpos($class, "\\")) {
            $prefix = substr($class, 0, $pos);
            if (isset(self::$prefixes[$prefix])) {
                $baseDir = self::$prefixes[$prefix];
            }
        }
        exit;

    }
}
