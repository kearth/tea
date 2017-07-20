<?php
namespace core;

class NullObject{
    public function __construct(){
        return new StdClass();
    }
}
