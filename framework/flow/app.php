<?php

namespace Tea\Framework\Flow;

class App extends Flow {

    protected string $key = 'app';

    private string $env;

    private string $action;

    private string $space;

    private const METHOD  = "method";

    private const DEFAULT_ENV = "dev";

    private const ENV = "env";

    private const NEXTTO = "nextTo";
    
    public function init(array $flow) : void {
        $flow = $this->getFlow($flow);
        $this->env = $flow[self::ENV] ?? self::DEFAULT_ENV;
        if (!isset($flow[self::NEXTTO]) || !isset($flow[self::METHOD]) || !method_exists($flow[self::NEXTTO], $flow[self::METHOD])) {
            throw new \Error("app conf error");
        }
        $this->nextTo($flow[self::NEXTTO], $flow[self::METHOD]);
    }

    public static function env() : string {
        return self::getInstance()->env;
    }

    public function nextTo(string $class, string $method){
        forward_static_call(array($class, $method));
    }

}

