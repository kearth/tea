<?php

namespace Tea\Core;

class SAPI
{

    public static $sapi = [
        "cgi-fcgi", 
        "cli", 
        "cli-server", 
        "continuity", 
        "embed", 
        "fpm-fcgi", 
        "isapi", 
        "litespeed", 
        "milter",
        "nsapi", 
        "phttpd", 
        "pi3web", 
        "roxen", 
        "thttpd", 
        "tux", 
        "webjames"
    ];

    public function start()
    {
        Config::init();
        Config::get("core", "init");
        Error::init();
        //日志
        //检测环境
        //预处理
        Request::init();
        //Response
        //provideRegister
        //dispatch
        //route
        //operate
        //exit
    }
}
