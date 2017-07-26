<?php
namespace BaseStone\Core;

class Response extends Core
{
    private static $instance = null;
    private $content = [
        'type'   => null,
        'params' => null
    ];

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

    public function getResponse()
    {
        return $this->content;
    }

    public function setResponse(String $type, Array $params)
    {
        $this->content['type']   = $type;
        $this->content['params'] = $params;
    }

}

