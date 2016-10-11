<?php
class ViewController{
    
    private $_viewDir = __DIR__.'/../View/';
    private $_baseViewDir = __DIR__.'/../View/html.php';
    private $_viewMoudle = array(
        __DIR__.'/../View/head.php'=>'<PHP-HEAD>',
        __DIR__.'/../View/style.php'=>'<PHP-STYLE>',
        __DIR__.'/../View/script.php'=>'<PHP-SCRIPT>',
        __DIR__.'/../View/body.php'=>'<PHP-BODY>'
    );

    private function getContent($viewDir){
        $content = file_get_contents($viewDir);
        return $content;
    }

    public function view($view,$data){
        $content = $this->getContent($this->_baseViewDir);
        foreach($this->_viewMoudle as $key => $value){
            $content = str_replace($value,$this->getContent($key),$content);
        }
        echo $content;
    }

}
