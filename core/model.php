<?php
namespace core;

class Model{

    public $id;
    public $createTime;
    public $updateTime;
    public $exist;
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
    
    protected function find($id){
        $sql = "select * from $this->tableName where id = $id";
        $res = $this->db->query($sql);
        return $res->fetchObject(get_class($this));
    }
    
    protected function getObjByParam($params){
        if(isset($params['columns']) && isset($params['conditions'])){
            $cols = $this->db->cols($params['columns']);
            $where = $this->db->where($params['conditions']);
            $sql = "select {$cols} from $this->tableName $where ";
            $res = $this->db->query($sql);
            return $res->fetchObject(get_class($this));
        }
    }

    public function isExist(){
        if($this instanceof NullObject){
            return false;
        }
        return true;
    }
}
