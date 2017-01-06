<?php
namespace controller;

class IndexController extends \core\BaseController{

    public function Test($requset,$response){
        $u = new \model\User();
        $i = $u->getUserById(1);
        //echo $u->create(array(
            //'account'=>'admin',
            //'password'=>'admin',
            //'status' => '0',
            //'nickname'=>'kunsama',
            //'telphone'=>'18888888888',
            //'email'=>'kunsama@163.com',
            //'headImg'=>'',
            //'tokens'=>''
        //));
       // $n = new DBTable('user',
       //     array(
       //         'id'=>'INT UNIQUE',
       //         'userName' =>'char(255)'
       //     ),
       //     array(
       //         'index1' =>'id' ,
       //         'index2' => 'userName',
       //     )
       // );
       // $db = DBConnect::getInstance();
       // $db->createTable($n);
    }

    public function index($requset,$response){
        print_r("hello world");
    }


    public function login($requset,$response){
        $data = array(
        );
        $response->setAttribute('view','index');
        $response->setAttribute('data',$data);
        $this->backToBrowser($response);   
    }

    public function verifyCode($requset,$response){
        $image = \lib\PicVerificationCode::create(100,30);
        echo $image;
    }

}

