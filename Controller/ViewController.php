<?php
class ViewController{
    
    private $_viewDir = __DIR__.'/../View/';
    private $_baseViewDir = __DIR__.'/../View/html.php';
    private $_view = null;
    private $_data = null;

    private function getContent($viewDir){
        $content = file_get_contents($viewDir);
        return $content;
    }

    public function view($view,$data){
        $this->_view = $this->_viewDir.$view.'.php';
        $this->_data = $data;
        $this->show();
    }

    private function show(){
        ob_start();
        $data = $this->_data;
        require $this->_view;
        ob_end_flush();
    }

    private function assembleData(){
    
    }



}
