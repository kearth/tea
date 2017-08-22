<?php
namespace Akf\Core;

class Bootstrap extends Singleton
{
    /**
     *  @name run 运行
     *  @access public
     *  @return void
     */
    public function run()
    {
        $modules = Config::getInstance()->getModules();
        foreach ($modules as $module) {
            if (class_exists($module)) {
                $module::run();
            } 
        }
    }
}

