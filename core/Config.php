<?php

namespace Tea\Core;

class Config
{
    use TeaTrait\Init;

    public static function load(string $fileName)
    {
        if (is_file($fileName)) {
            require($fileName);
        } else {
            throw new \Exception("文件路径错误", 1);
        }
    }
}

