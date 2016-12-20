<?php
namespace core;
class Request{

    public $params;
    public function __construct(){
        $this->params = $_REQUEST;
    }
}
