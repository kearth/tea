<?php 

namespace Tea\Framework;

/**
 * 引导类
 */
class Bootstrap {

    private const FLOWLIST = "flowlist";

    private const FLOW = "flow";

    private const PLUGINLIST = "pluginlist";

    private const PLUGIN = "plugin";

    private const CLASSNAME = "class";

    private const METHOD  = "method";

    private static array $conf = array();

    private static string $confDir = "/conf/";


    /**
     * Run 
     */
    public static function run(string $rootPath) {
        self::init($rootPath);
        self::startFlow($rootPath);
    }

    private static function init(string $rootPath) : void {
        $dir = $rootPath . self::$confDir;
        if (!is_dir($dir)){
            throw new Error("conf dir is not exist!");
        }
        if ($handle = opendir($dir)) {
            while (false !== ($file = readdir($handle))) {
                if ("ini" == pathinfo($file, PATHINFO_EXTENSION)) {
                    $conf = parse_ini_file($dir . $file, true);
                    if ($conf){
                        self::$conf = array_merge(self::$conf, $conf);
                    }
                }
            }
            closedir($handle);
        }
    }

    private static function startFlow(string $rootPath) {
        require($rootPath . "/framework/autoload.php");
        Autoload::init($rootPath);
        $pluginList = self::$conf[self::PLUGINLIST][self::PLUGIN] ?? array();
        foreach ($pluginList as $k => $plugin){
            $plugin = self::$conf[$plugin] ?? array();
            if (is_array($plugin)) {
                $class = $plugin[self::CLASSNAME] ?? "";
                $method = $plugin[self::METHOD] ?? "";
                Register::add($class, new $class());
            }
        }
        $flowList = self::$conf[self::FLOWLIST][self::FLOW] ?? array();
        foreach ($flowList as $k => $flowname){
            $flow = self::$conf[$flowname] ?? array();
            if (is_array($flow)) {
                $class = $flow[self::CLASSNAME] ?? "";
                $method = $flow[self::METHOD] ?? "";
                $instance = new $class();
                $instance->$method(self::$conf);
                Register::add($class, $instance);
            }
        }
    }

}
