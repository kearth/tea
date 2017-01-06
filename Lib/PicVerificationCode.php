<?php
namespace lib;

class PicVerificationCode{

    const colors  = array(
       "black"  =>array(0,0,0),
       "white"  =>array(255,255,255),
       "gray"   =>array(190,190,190),
       "red"    =>array(255,0,0),
       "green"  =>array(0,255,0),
       "blue"   =>array(0,0,255),
       "yellow" =>array(255,255,0),
       "orange" =>array(255,165,0),
       "purple" =>array(160,32,240),
       "pink"   =>array(255,181,197),
       "cyan"   =>array(0,255,255),
       "Magenta"=>array(255,0,255),
    );
    const charset = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'; 

    const colorset = array(
        "gray","red","green","blue","yellow","orange","purple","pink","cyan","Magenta"
    );

    const ttf = "/development/git/ksmCMS/Storage/锐字逼格锐线体简4.0.ttf";

    public static function create($width,$height){
        $codeLen = 4;//验证码长度
        header("Content-type:image/png");
        $img=imagecreatetruecolor($width,$height);
        $bgColor = self::getColor($img,'white');
        imagefill($img,0,0,$bgColor);
        self::getCode($img,$codeLen,$width,$height);
        imagepng($img);
        imagedestroy($img);
    }

    private static function getCode($img,$length,$width,$height){
        session_start();
        $session = "";
        $scope = strlen(self::charset);
        $colorScope = sizeof(self::colorset);
        $charsetArray = str_split(self::charset);
        $x = $width/5;
        $y = $height/2;
        for($i = 0;$i<$length;$i++){
            $char= $charsetArray[rand(0,$scope-1)];
            $session .= $char; 
            $color = self::colorset[rand(0,$colorScope-1)];
            imagefttext($img,20,(8-rand(0,16))*10,$x*(1+$i),$y,self::getColor($img,$color),self::ttf,$char);
        }
        $_SESSION['code'] = $session;

    }

    private static function getColor($img,$color){
        $R = self::colors[$color][0]; 
        $G = self::colors[$color][1]; 
        $B = self::colors[$color][2]; 
        return imagecolorallocate($img,$R,$G,$B);
    }
    

}
