<?php

class Person {    
	/**  
	 * For the sake of demonstration, we"re setting this private 
	 */   
	private $_allowDynamicAttributes = false;  

	/** type=primary_autoincrement */  
	protected $id = 0;  

	/** type=varchar length=255 null */  
	protected $name;  

	/** type=text null */  
	protected $biography;  

	public function __construct(){

	}


	public function getId()  
	{  
		return $this->id;  
	}  
	public function setId($v)  
	{  
		$this->id = $v;  
	}  
	public function getName()  
	{  
		return $this->name;  
	}  
	public function setName($v)  
	{  
		$this->name = $v;  
	}  
	public function getBiography()  
	{  
		return $this->biography;  
	}  
	public function setBiography($v)  
	{  
		$this->biography = $v;  
	}  
}  



$r = exec("./hello");
echo $r;



/*
$class = new ReflectionClass('Person');
$instance = $class->newInstanceArgs(array($args));
var_dump($instance);


function replace_unicode_escape_sequence($match) {
return mb_convert_encoding(pack('H*', $match[1]), 'UTF-8', 'UCS-2BE');
}
$name = '\u4f63\u91d1\u603b\u91d1\u989d';
$str = preg_replace_callback('/\\\\u([0-9a-f]{4})/i', 'replace_unicode_escape_sequence', $name);
echo $str; //输出： 新浪微博

*/
