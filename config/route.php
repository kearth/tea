<?php
return [
    '{controller|[a-zA-Z]+}/{action}' => function ($controller, $action) {
        return ucfirst($controller) . "/" . $action;
    }
];
