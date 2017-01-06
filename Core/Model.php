<?php
namespace core;

class Model{

    protected $id;
    protected $createTime;
    protected $updateTime;
    protected $exist;
    protected $tableName;
    protected $db;

    public function __construct(){
        $this->db = DBConnect::getInstance();
        if(empty($this->tableName)){
            $this->tableName = basename(str_replace('\\','/',get_class($this)));
        }
    }

    protected function create($className){
        $class = new \ReflectionClass($className);
        $col = "insert into ".$this->tableName."(";
        $value = ") values(";
        foreach($class->getProperties(\ReflectionProperty::IS_PUBLIC) as $prop)
        {
            $propName = $prop->getName();
            if(empty($this->$propName)){
                continue;
            }
            $col     .= $propName.",";
            $value  .= "'".$this->$propName."',";
        }
        $sql = rtrim($col,',').rtrim($value,',').");";
        return $this->db->insert($sql);
    }
    
    protected function getObjById($id){
        $sql = "select * from $this->tableName where id = $id";
        $res = $this->db->query($sql);
        return $res->fetchObject(get_class($this));
    }

}
