<?php
namespace BaseStone\Core;

class Request extends Core
{
 
    private static $instance = null;

    private function __construct()
    {
        $this->http_user_agent = $_SERVER['HTTP_USER_AGENT'];
        $this->server_protocal = $_SERVER['SERVER_PROTOCOL'];
        $this->request_uri     = $_SERVER['REQUEST_URI'];
        $this->request_method  = $_SERVER['REQUEST_METHOD'];
        $this->content_length  = $_SERVER['CONTENT_LENGTH'];
        $this->content_tyep    = $_SERVER['CONTENT_TYPE'];
        $this->cookie          = $_COOKIE;
        $this->get             = $_GET;
        $this->post            = $_POST;
        $this->env             = $_ENV;
        $this->action          = $_REQUEST['rewrite'];
        unset($_REQUEST['rewrite']);
        $this->params          = $_REQUEST;
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
       
}

