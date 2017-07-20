<?php
class A 
{

    public function __construct($a)
    {
        echo $a."\n";
    }

    public function __construct1($a,$b)
    {
        echo $a.$b."\n";
    }

    public function __construct2($a,$b,$c)
    {
        echo $a.$b.$c."\n";
    }
}

new A("1");
new A("2","33");
new A("33","2","1");
