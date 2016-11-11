 <?php
 class IndexController extends ViewController{
     
     public function Test(){
        $data = array(
            'nav'=>array(
                'logoName'=>'KunCMS',
                'userStatus'=>'Logout',
                'userSet'=>'Settings',
                'userName'=>'Hi,eric'
            ),
            'content'=>array(
                
            )
        );
        $this->view("html",$data);
     }


 }

