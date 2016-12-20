    <?php require "nav.php";?>
    <?php
        if(array_key_exists('content',$data)){
            echo "<div class='content'>
                <div class='loginTable'>
                    <form>
                    <lable>account:</lable><input type='text'>
                    <br/>
                    <lable>password:</lable><input type='text'>
                    </form>
                </div>
                  </div>";
        }
    ?>


    <?php// require "left.php";?>
    <?php// require "right.php";?>
