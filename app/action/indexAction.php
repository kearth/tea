<?php 

namespace App\Action;

class IndexAction extends BaseAction {

    public static function execute() {
        $output = self::apiFormat(array("welcome to Tea framework"));
        self::response($output);
    }

}
