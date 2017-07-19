<?php

return array(
    'DBDefault' => array(
        'dsn'  => 'mysql:host=127.0.0.1;dbname=mysql;',
        'user' => 'root',
        'password' => '=lV0ia3u#d*z',
        'option' => array(
            \PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES'utf8';"
        ),
    ),
    //'MySql' => array(
        //'dsn'  => 'mysql:host=127.0.0.1;dbname=mysql;',
        //'user' => 'root',
        //'password' => '123456',
        //'option' => array(
            //\PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES'utf8';"
        //),
    //),
    'Log' => array(
        'path' => ROOT."/storage/",
        'file' => 'error'
    )
);
