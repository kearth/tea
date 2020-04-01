<?php 

/**
 * 异常配置
 */

// 错误码 
define('CODE_SUCCESS', 0);
define('CODE_UNKNOWN', 999);
define('CODE_FRAMEWORK', 1);

// 配置项
return array(
    'errMsgMap' => array(
       0 => 'success', 
       1 => 'framework error', 
       999 => 'unknown error', 
    ) 
);
