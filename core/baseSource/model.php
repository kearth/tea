<?php

namespace Akf\Core\BaseSource;

abstract class Model
{
    private $handle;

    protected function __construct(DataSource $dataSource)
    {
        $this->handle = $dataSource;
    }

    public function get($name = '')
    {
        if (empty($name)) {
            return $this->prototype;
        } else {
            return $this->prototype[$name];
        }
    }
}
