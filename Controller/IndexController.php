<?php
namespace controller;

class IndexController extends \core\BaseController{

    public function Test(){
        var_export($this->request);
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
        $this->view("html",$data);
    }

    public function index(){
        print_r("hello world");
    }

}

