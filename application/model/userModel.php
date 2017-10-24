<?php

class UserModel extends Model
{
    public $prototype = [
        'id' => 1,
        'name',
        'ctime',
        'utime',
        'ver'
    ];

    public function __construct()
    {
        //$db = Container::make('DB');
        //parent::__construct($db);
    }


}
