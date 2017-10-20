<?php
return [
    '{controller|[a-zA-Z]+}/{action}' => function ($controller, $action) {
        return $controller . "Controller" . "/" . $action . "Action";
    }
];
