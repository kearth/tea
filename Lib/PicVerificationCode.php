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

    public static function create($width,$height){
        $codeLen = 4;//验证码长度
        header("Content-type:image/png");
        $img=imagecreatetruecolor($width,$height);
        $bgColor = self::getColor($img,'white');
        imagefill($img,0,0,$bgColor);
        self::getCode($img,$codeLen);
        imagepng($img);
        imagedestroy($img);
    }

    private static function getCode($img,$length){
        $scope = strlen(self::charset);
        $colorScope = sizeof(self::colorset);
        $charsetArray = str_split(self::charset);
        $x = 20;
        $y = 15;
        for($i = 0;$i<$length;$i++){
            $char= $charsetArray[rand(0,$scope-1)];
            $color = self::colorset[rand(0,$colorScope-1)];
            imagechar($img,5,$x+10*$i,$y,$char,self::getColor($img,$color));
        }
    }

    private static function getColor($img,$color){
        $R = self::colors[$color][0]; 
        $G = self::colors[$color][1]; 
        $B = self::colors[$color][2]; 
        return imagecolorallocate($img,$R,$G,$B);
    }
    

}
