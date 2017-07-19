<?php
namespace Bootstrap;

use Bootstrap\Autoload;
include_once(ROOT."/bootstrap/autoload.php");

class Bootstrap
{
    private static $instance;

    private function __construct()
    {
    }

    private function __clone()
    {
    
    }

    public static function getInstance()
    {
        if (!(self::$instance instanceof Bootstrap)) {
            self::$instance = new static();  
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

