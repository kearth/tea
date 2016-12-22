<?php
namespace core;

class DBConnect extends \PDO{
    public function __construct($dsn="mysql:host=127.0.0.1;dbname=ksm_dev;",$user="root",$password="123456",$option=array(\PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES'utf8';")){
        try{
            parent::__construct($dsn,$user,$password,$option);
        } catch (\PDOException $e){
            echo 'Connection failed:' . $e->getMessage();
        }
    }

    public function testConn(){
        $sql = "select * from connect_test";
        foreach($this->query($sql) as $row){
            print $row['id']."\n";
            print $row['name']."\n";
            print $row['result']."\n";
        }
        return $result;
    }

}
