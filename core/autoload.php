<?php
namespace Akf\Core;

class Autoload
{
    private $hasRequiredFile = [];

    public static function register()
    {
        spl_autoload_register(array(new static(),'loadClass'));
    }

    public function loadClass($class)
    {
        if (file_exists(CONTROLLER . $class . '.php')) {
            $classFile = CONTROLLER . $class . '.php';
        } elseif (file_exists(MODEL . $class . '.php')) {
            $classFile = MODEL . $class . '.php';
        } else {
            $classFile = ROOT_PATH . DIRECTORY_SEPARATOR . substr(str_replace('\\', '/', $class), 4) . '.php';
        }

        if (in_array($classFile,$this->hasRequiredFile)) {
            return;
        }

        if (file_exists($classFile)) {
            //echo $classFile . "<br>";
            require $classFile;
            $this->has_required_file[] = $classFile;
        }
    }
}

