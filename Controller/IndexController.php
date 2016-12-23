<?php
namespace controller;

class IndexController extends \core\BaseController{

    public function Test($requset,$response){
        $data = array(
            'nav'=>array(
                'logoName'=>'KunCMS',
                'userStatus'=>'Logout',
                'userSet'=>'Settings',
                'userName'=>'Hi,eric'
            ),
            'content'=>array(

            )
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
            'nav'=>array(
                'logoName'=>'KunCMS',
                'userStatus'=>'Logout',
                'userSet'=>'Settings',
                'userName'=>'Hi,eric'
            ),
            'content'=>array(

            )
        );
        $response->setAttribute('view','index');
        $response->setAttribute('data',$data);
        $this->backToBrowser($response);   
    }

}

