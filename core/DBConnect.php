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

    public function insert($sql){
        return static::$_conn->exec($sql);
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
        return static::$_conn->exec($sql);
    }
    
    public function where($condition){
        $where = "where ";
        if (is_array($condition)) {
            foreach($condition as $k => $v){
                $where .= " {$k} = '{$v}' and ";
            }
            $where .= " 1=1 ";
        } else if(is_string($condition)) {
            $where .= $condition;
        } else {
            $where .= " 1=1 ";
        }
        return $where;
    }

    public function cols($condition){
        $cols = "";
        if(is_string($condition)) {
            $cols = $condition;
        } else {
            $cols = " * ";
        }
        return $cols;
    }
}
