<?php

define("AUTOLOAD_ALIAS", "alias");
define("AUTOLOAD_NAMESPACE", "namespace");

return [
     AUTOLOAD_ALIAS => [
         "Tea/kernel/Container" => "Container"

     ],
     AUTOLOAD_NAMESPACE => [
         "Tea" => ROOT_PATH . "/"
     ],
];
