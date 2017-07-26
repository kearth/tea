<?php
namespace BaseStone\Bootstrap;

class Autoload
{
    private static $instance = null;
    private $has_required_file = [];

    private function __construct()
    {
    }
    
    private function __clone()
    {
    }

    private function __sleep()
    {
    }

    private function __wakeup()
    {
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    public function register()
    {
        spl_autoload_register(array($this,'loadClass'));
    }

    public function loadClass($class)
    {
        $class_file = ROOT_PATH . DIRECTORY_SEPARATOR . str_replace('\\', '/', $class) . '.php';
        
        if (in_array($class_file,$this->has_required_file)) {
            return;
        }

        //echo $class_file."<br/>";
        if (file_exists($class_file)) {
            require $class_file;
            $this->has_required_file[] = $class_file;
        }
    }
    
    /** 遍历指定目录 **/
    //public function readAllDir($dir)
    //{
        //$dir_tree = [];
        //$current_dir = scandir($dir);
        //foreach ($current_dir as $key => $value) {
            //if (!in_array($value,['.','..'])) {
                //if (is_dir($dir . DIRECTORY_SEPARATOR .$value)) {
                    //$dir_tree[$value] = $this->readAllDir($dir .DIRECTORY_SEPARATOR . $value);
                //} else {
                    //if (substr($value, -4) === '.php') {
                        //$dir_tree[] = $value;
                    //}
                //}
            //}
        //}
        //return $dir_tree;
    //}

    /** 整理类文件数组 **/
    //public function arrangeClassLib($dir, $class_lib)
    //{
        //$new_class_lib = [];
        //foreach($class_lib as $key => $value) {
            //if (is_array($value)) {
                //$new_class_lib = array_merge($new_class_lib, $this->arrangeClassLib($dir. DIRECTORY_SEPARATOR .$key, $value));
            //} else {
                //if (isset($new_class_lib[strtolower($value)])) {
                    //$new_class_lib[strtolower($value)][] = $dir . DIRECTORY_SEPARATOR .$value;
                //} else {
                    //$new_class_lib[strtolower($value)] = $dir . DIRECTORY_SEPARATOR .$value;
                //}
            //}
        //}
        //return $new_class_lib;
    //}
}

