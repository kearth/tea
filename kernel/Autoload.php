<?php

namespace Tea\Kernel;

/**
 *
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

    protected static function loadMappedFile($prefix, $baseDir, $class)
    {
        $path = $baseDir . "/" .  str_replace("\\", "/", $class) . ".php";
        self::requireFile($path);
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
            $clazz = substr($class, $pos + 1);
            if (isset(self::$prefixes[$prefix])) {
                $baseDir = self::$prefixes[$prefix];
                self::loadMappedFile($prefix, $baseDir, $clazz);
            }

        }
    }
}
