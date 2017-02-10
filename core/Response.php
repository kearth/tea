<?php
namespace core;

class Response{

    private $_view;
    private $_data;

    public function __construct(){
    }

    public function setAttribute($attr,$value){
        $attrName = '_'.$attr;
        $this->$attrName = $value;
    }

    public function getAttribute($attr = 'all'){
        if($attr === 'all'){
            return $this;
        } else {
            $attrName = '_'.$attr;
            return $this->$attrName;
        }
    }

}
