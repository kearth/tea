<?php
namespace core;

class Template{

    private $_template;
    private $_data;

    public function __construct($template,$data){
        $this->_template = file_get_contents($template);
        $this->_data     = $data;    
    }

    public function template2html(){
        $this->replace();
        $this->output();
    }

    private function replace(){
        foreach($this->_data as $var => $val){
            $find = '{$'.$var.'}';
            $this->_template =  str_replace($find,$val,$this->_template);
        }
    }

    private function output(){
        echo $this->_template;
    }



}
