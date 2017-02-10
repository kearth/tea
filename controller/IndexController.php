<?php
namespace controller;

class IndexController extends \core\BaseController{

    public function isLogin(){
        if(isset($_COOKIE['uid'])){
            $uid = $_COOKIE['uid'];
            session_start();
            if(isset($_SESSION[$uid])){
                return true;
            }
        }
        return false;
    }

    public function index($requset,$response){
        if($this->isLogin()){
            $data = [
                'userStatus' => 'sign out'
            ];
            $response->setAttribute('view','index');
            $response->setAttribute('data',$data);
            $this->backToBrowser($response);
        } else {
            $this->login($requset,$response);
        }
    }


    public function login($requset,$response){
        $data = array(
        );
        $response->setAttribute('view','login');
        $response->setAttribute('data',$data);
        $this->backToBrowser($response);   
    }

    public function signin($requset,$response){
        session_start();
        $result = array();
        $account    = $requset->params['account'];
        $password   = $requset->params['password'];
        $verifyCode = $requset->params['verifyCode'];
        
        if(empty($verifyCode)){
            $result['retData'] = "请输入验证码！";
        } elseif(strtolower($verifyCode) != strtolower($_SESSION['code'])){
            $result["retData"] = "您输入的验证码有误，请重新输入！";
        } else {
            if(empty($account)){
                $result["retData"] = "账号不能为空！";
            } elseif(empty($password)) {
                $result["retData"] = "密码不能为空！";
            } else{
                $user = new \model\User();
                $user = $user->getUserBeforeLogin($account);
                if($user->isExist()){
                    if($password === $user->password){
                        $uid = md5($user->id.time());
                        setcookie("uid",$uid,time()+3600);
                        $_SESSION[$uid]  = $account;
                        $result["retData"] = "登录成功!";
                        $result["url"] = "index";
                    } else {
                        $result["retData"] = "密码错误!";
                    }
                } else {
                    $result["retData"] = "该用户不存在!";
                }
            }
        }
        echo json_encode($result);
    }

    public function verifyCode($requset,$response){
        echo \lib\PicVerificationCode::create(100,30);
    }

}

