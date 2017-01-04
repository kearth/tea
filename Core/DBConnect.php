<?php
namespace core;

class DBConnect{

    protected static $_instance = null;
    protected static $_conn = null;

    protected function __construct($dbConfig){
        $config = Config::getConfig($dbConfig);
        try{
            static::$_conn = new \PDO($config['dsn'],$config['user'],$config['password'],$config['option']);
        } catch (\PDOException $e){
            echo 'Connection failed:' . $e->getMessage();
        }
    }

    protected function __clone(){
    
    }

    public static function getInstance($dbConfig = "DBDefault" ){
        //单次调用，拒绝跨库
        if(static::$_instance === null){
            static::$_instance = new static($dbConfig);
        }
        return static::$_instance;
    }

    public function query($sql){
        return static::$_conn->query($sql);
    }

    public function createDataBase($dbName){
        $sql = "CREATE DATABASE IF NOT EXISTS $dbName DEFAULT CHARSET utf8 COLLATE utf8_general_ci;";
        return static::$_conn->exec($sql);
    }
    
    public function removeDataBase($dbName){
        $sql = "DROP DATABASE IF EXISTS $dbName;";
        return static::$_conn->exec($sql);
    }
    
    public function createTable(DBTable $table){
        $sql = $table->getCreateSQL();
        echo $sql;
        return static::$_conn->exec($sql);
    }



    public function testConn(){
        $sql = "select id from connect_test;";
        $res = $this->query($sql);
        foreach($res as $row){
            print_r($row['id']);
        }
        return $res;
    }
}
