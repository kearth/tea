<?php
namespace core;

class View{
    private $_viewDir = VIEW."/";
    private $_view;
    private $_data;

    public function __construct($view,$data){
        $this->_view = $this->_viewDir.$view.'.html';
        $this->_data = $data;
    }

    public function show(){
        ob_start();
        $data = $this->_data;
        include($this->_view);
        ob_end_flush();
    }

}
