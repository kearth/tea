<?php
namespace controller;

class IndexController extends \core\BaseController{

    public function Test($requset,$response){
        $data = array(
        );
        $response->setAttribute('view','index');
        $response->setAttribute('data',$data);
        $this->backToBrowser($response);   
    }

    public function index($requset,$response){
        print_r("hello world");
    }


    public function login($requset,$response){
        $data = array(
        );
        $response->setAttribute('view','index');
        $response->setAttribute('data',$data);
        \lib\PicVerificationCode::create(100,40);
        //$this->backToBrowser($response);   
    }

}

