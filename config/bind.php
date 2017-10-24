<?php

Container::bind('Api', function () {
    return new Api();
});

Container::bind('View', function () {
    return new View();
});

Container::bind('Stream', function ($request) {
    return new Stream($request);
});

Container::bind('UserModel', function () {
    return new UserModel();
});

Container::bind('PostController', function () {
    return new PostController();
});

