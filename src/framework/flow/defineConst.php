<?php 

namespace Tea\Framework\Flow;

class DefineConst extends Flow {

    protected string $key = "const";

    public function init(array $flow) {
        $flow = $this->getFlow($flow);
        foreach($flow as $k => $v){
            if (!defined($k)){
                define($k, $v);
            }
        }
    }
}
