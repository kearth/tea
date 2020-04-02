<?php

namespace Tea\Framework;

/**
 * 分发类
 */
class Dispatcher {

    // 默认Controller
    const DEFAULT_CONTROLLER = 'Index';
    // 默认Action
    const DEFAULT_ACTION = 'index';

    /**
     * 分发
     */
    public static function dispatch() : void {
        $router     = Router::getRouter();
        $routerMap  = explode('_', $router);
        $controller = $routerMap[0] ?? static::DEFAULT_CONTROLLER;
        $action     = $routerMap[1] ?? static::DEFAULT_ACTION;
        $controller = static::getController($controller);
        $controllerObj = new $controller();
        if (!method_exists($controllerObj, $action)) {
            Error::throw("{$action}方法不存在");
        }
        $controllerObj->$action();
    }

    /**
     * 获取Controller
     */
    private static function getController(string $controller) : string {
        $controllerMap = explode('/', $controller);
        $newController = Config::get('ControllerNamespace');
        foreach($controllerMap as $controllerPart) {
            $newController .= ucfirst($controllerPart) . '\\';
        }
        $newController = rtrim($newController, '\\') . 'Controller';
        return $newController;
    }

}

