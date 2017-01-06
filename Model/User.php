<?php
namespace model;

class User extends \core\Model{

    public $account;    //账户
    public $password;   //密码
    public $status;     //0未登录,1登录,2离线
    public $nickname;   //昵称
    public $telphone;   //手机号
    public $email;      //邮箱
    public $headImg;    //头像
    public $tokens;     //其他登录方式权限凭证

    public function __construct(){
        $this->tableName = 'User';
        parent::__construct();
    }

    public function create($user){
        $this->account   = $user['account'];
        $this->password  = $user['password'];
        $this->nickname  = $user['nickname'];
        $this->status    = $user['status'];
        $this->telphone  = $user['telphone'];
        $this->email     = $user['email'];
        $this->headImg   = $user['headImg'];
        $this->tokens    = $user['tokens'];
        parent::create(__CLASS__);
    }
    
    public function getUserById($id){
        return parent::getObjById($id);
    }

}
