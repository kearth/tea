<?php
namespace Bootstrap;

use Bootstrap\Autoload;

class Bootstrap
{
    private static $instance = null;

    /** 私有构造方法 **/
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

    public function run()
    {
        $this->init();
        $this->load();
        $this->work();
        $this->over();   
    } 

    private function init()
    {
        $autoload = new Autoload();
        $autoload->register();
    }
    
    private function load()
    {

    }

    private function work()
    {
    
    }

    private function over()
    {
    
    }
}

