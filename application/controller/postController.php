<?php

class PostController
{
    public function indexAction()
    {
        
        $this->response = [
            'type' => 'View',
            'code' => 1,
            'msg'  => 2,
            'data' => [22332]
        ];
    }
}
