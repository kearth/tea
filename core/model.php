<?php

namespace Akf\Core;

abstract class Model
{
    private $handle;

    protected function __construct(DataSource $dataSource)
    {
        $this->handle = $dataSource;
    }

    //public function create()
    //{
        //return $this->handle->create();   
    //}

    //public function save()
    //{
        //return $this->handle->save();
    //}

    //public function remove()
    //{
        //return $this->handle->remove();
    //}

    //public function find(array $param)
    //{
        //return $this->handle->find($param);
    //}

    public function get($name = '')
    {
        if (empty($name)) {
            return $this->prototype;
        } else {
            return $this->prototype[$name];
        }
    }
}
