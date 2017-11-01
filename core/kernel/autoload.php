<?php
namespace Akf\Core\Kernel;

class Autoload
{
    private static $hasRequiredFile = [];
    private static $aliasClass = [];

    public static function register(array $aliasConfig)
    {
        self::$aliasClass = $aliasConfig;
        spl_autoload_register(array(new static(),'loadClass'));
    }

    public function loadClass($class)
    {
        self::setAlias($class);
        if (file_exists(CONTROLLER . $class . '.php')) {
            $classFile = CONTROLLER . $class . '.php';
        } elseif (file_exists(MODEL . $class . '.php')) {
            $classFile = MODEL . $class . '.php';
        } else {
            $classFile = ROOT_PATH . DIRECTORY_SEPARATOR . substr(str_replace('\\', '/', $class), 4) . '.php';
        }

        if (in_array($classFile, self::$hasRequiredFile)) {
            return;
        }

        if (file_exists($classFile)) {
            //echo $classFile . "<br>";
            require $classFile;
            self::$hasRequiredFile[] = $classFile;
        }
    }

    public static function setAlias(string $name)
    {   
        $class = array_search($name, self::$aliasClass);
        if ($class) {
            class_alias($class, $name);
        }
    }

}

