<?php
namespace BaseStone\Core;

class Response extends Core
{
    private static $instance = null;
    private $params = [];

    private function __construct()
    {

    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    public function getParams(): array
    {
        return $this->params;
    }

    public function setParams(array $params)
    {
        $this->params = $params;
    }

}

