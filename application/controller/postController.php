<?php

class PostController
{
    public function indexAction()
    {
        //$userModel = Container::make('UserModel');
        //echo $userModel->get('id');
       //exit;
        $this->response = [
            'type' => 'View',
            'code' => 1,
            'msg'  => 2,
            'data' => [22332]
        ];
    }
}
