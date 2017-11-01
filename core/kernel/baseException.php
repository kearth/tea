<?php

namespace Akf\Core\Kernel;

/**
 *  基础异常
 */
class BaseException extends \Exception
{
    public function __toString()
    {   
        return get_class($this) 
             . " Error Info : "
             . $this->getMessage()
             . " , Error Code : "
             . $this->getCode()
             . " in "
             . $this->getFile()
             . " on Line "
             . $this->getLine();
    }
}

