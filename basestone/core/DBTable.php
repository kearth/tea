<?php
namespace core;

class DBTable{
    public $tableName;
    public $column;
    public $index;

    public function __construct($tableName,array $column,array $index){
        $this->tableName = $tableName;
        $this->column = $column;
        $this->index = $index;
    }


    public function getCreateSQL(){
        $sql = "CREATE TABLE $this->tableName (";
        foreach($this->column as $k => $v){
            $sql .= "$k $v,";
        }
        foreach($this->index as $k => $v){
            $sql .= "INDEX $k (".$v."),";
        }
        $sql = rtrim($sql,',');
        $sql .= ");"; 
        return $sql;
    }

}
