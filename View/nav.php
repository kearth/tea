<div class="nav">
<?php
    if(array_key_exists('nav',$data)){
        foreach($data['nav'] as $key => $value){
            echo "<span id=\"{$key}\">{$value}</span>";
        }
    }
?>
</div>
