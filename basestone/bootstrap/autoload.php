<?php
namespace BaseStone\Bootstrap;

class Autoload
{
    private $class_lib = null;

    public function __construct()
    {
        $this->class_lib = $this->arrangeClassLib(ROOT_PATH."/basestone",$this->readAllDir(ROOT_PATH."/basestone/"));
    }
    
    public function register()
    {
        spl_autoload_register(array($this,'loadClass'));
    }

    public function loadClass($class)
    {
        $class_file_name = strtolower($class).".php";
        if (!empty($this->class_lib) && array_key_exists($class_file_name,$this->class_lib)) {
            if (file_exists($this->class_lib[$class_file_name])) {
                var_dump($this->class_lib[$class_file_name]);
                include($this->class_lib[$class_file_name]);
            }
        }
    }
    
    /** 遍历指定目录 **/
    public function readAllDir($dir)
    {
        $dir_tree = [];
        $current_dir = scandir($dir);
        foreach ($current_dir as $key => $value) {
            if (!in_array($value,['.','..'])) {
                if (is_dir($dir . DIRECTORY_SEPARATOR .$value)) {
                    $dir_tree[$value] = $this->readAllDir($dir .DIRECTORY_SEPARATOR . $value);
                } else {
                    if (substr($value, -4) === '.php') {
                        $dir_tree[] = $value;
                    }
                }
            }
        }
        return $dir_tree;
    }

    /** 整理类文件数组 **/
    public function arrangeClassLib($dir, $class_lib)
    {
        $new_class_lib = [];
        foreach($class_lib as $key => $value) {
            if (is_array($value)) {
                $new_class_lib = array_merge($new_class_lib, $this->arrangeClassLib($dir. DIRECTORY_SEPARATOR .$key, $value));
            } else {
                if (isset($new_class_lib[strtolower($value)])) {
                    $new_class_lib[strtolower($value)][] = $dir . DIRECTORY_SEPARATOR .$value;
                } else {
                    $new_class_lib[strtolower($value)] = $dir . DIRECTORY_SEPARATOR .$value;
                }
            }
        }
        return $new_class_lib;
    }

}

