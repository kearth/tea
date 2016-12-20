<?php
namespace core;

class BaseController{

    private $_viewDir = VIEW."/";
    private $_baseViewDir = VIEW.'/html.php';
    private $_view = null;
    private $_data = null;
    private $_assembleData = array('title','include','style','script','body');
    protected $request;
    protected $response;

    public function __construct($request,$response){
        $this->request = $request;
        $this->response = $response;
    }

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

    private function assembleData($title,$include,$style,$script,$body){
        $data = array();
        foreach($this->_assembleData as $key){
             $data[$key] = $$key;
        }
        return json_encode($data);
    }
}
