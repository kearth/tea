<?php
namespace core;

class View{
    private $_viewDir = VIEW."/public/";
    private $_view;
    private $_data;
    private $_mode;
    const TEMPLATE = 1;

    public function __construct($view,$data){
        $this->_mode = Config::getConfig('ViewMode',CONFDIR."/coreConf.php");
        $this->_view = $this->_viewDir.$view.'.html';
        $this->_data = $data;
    }

    public function show(){
        ob_start();
        if($this->_mode == self::TEMPLATE){
            $template = new Template($this->_view,$this->_data);
            $template->template2html();
        } else {
            $data = $this->_data;
            include($this->_view);
        }
        ob_end_flush();
    }


}
