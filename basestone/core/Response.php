<?php
namespace BaseStone\Core;

class Response extends Singleton
{
    private $outputType;
    private $params = [];

    public function getParams(): array
    {
        return $this->params;
    }

    public function setParams(array $params)
    {
        $this->params = $params;
    }

    public function getOutput()
    {
        return $this->outputType;
    }

    public function setOutput(string $outputType)
    {
        $this->outputType = $outputType;
    }
}

