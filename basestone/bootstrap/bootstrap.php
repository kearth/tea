<?php
namespace BaseStone\Bootstrap;

class Bootstrap
{
    private static $instance = null;

    private function __construct()
    {
    
    }

    private function __clone()
    {
    
    }

    private function __wakeup()
    {
    
    }

    private function __sleep()
    {
    
    }

    public static function getInstance()
    {
        if (null === self::$instance) {
            self::$instance = new self();
        }
        return self::$instance;
    }
    public function run()
    {
        $this->init();
        $this->work();
        $this->over();   
    } 

    /** 初始化 **/
    private function init()
    {
        Config::getInstance()->load(CONFIG_PATH);

        //日志系统
        
        //错误处理系统
    }

    /** 内容分发处理 **/
    private function work()
    {
        //路由
        $router = Router::getInstance();
        $request = $router->getRequest();
        var_export($request);


        //分发工作
        //Dispath
    }

    /** 结束 **/
    private function over()
    {
        //Response
        
        //清理结果
    }
}

