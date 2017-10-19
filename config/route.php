<?php
return [
    'post/{name|[a-z]}/post/{id}' => function($name) {
        return "33" . $name[0];
    },
    'post/{name}' => function($uri) {
        return '88833/' . $uri;
    },
    'foo' => function() {
        return '3333';
    }

];
