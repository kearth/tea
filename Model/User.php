<?php
namespace model;

class User extends \core\Model{

    public $account;
    public $password;
    public $status;
    public $nickname;
    public $telphone;
    public $email;
    public $headImg;
    public $tokens;

    public function __construct(){
        $this->exist = 4;
    }


    public function create(){
        
    }

    public function getUserById(){}

}
