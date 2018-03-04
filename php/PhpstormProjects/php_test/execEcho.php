<?php
exec("ls -l", $output, $error);
//exec('wkhtmltopdf http://somesite.com /home/user/file.pdf');
echo "<pre>";
print_r( $output );
?>

