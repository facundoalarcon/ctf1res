<?php
error_reporting(E_ALL & ~E_WARNING & ~E_NOTICE);
$FLAG = "flag{Y0u_ll_n0t_g377ing_1t_fR0m_h3r3}";
$current_time = time();
$TOKEN = $current_time.rand(1, 30);
if(isset($_POST['SubmitButton'])){ //check if form was submitted
  $username = $_POST['username']; //get input text
  $in_token = $_POST['token']; //get input text

  if ($in_token === $TOKEN && $username.$in_token.rand(1,31337)=="0e12021810133773311610")
		$msg = "<h1 class=\"error\"> Well done! <br /> $FLAG </h1> <br><br>";
  else
		$msg = "<h1 class=\"error\"> Wrong token </h1> <br><br>";
}
?>

<html>
<head>
<link rel="icon" href="/favicon.ico" type="image/x-icon" >
  <meta charset="UTF-8">
  <title>SuperSecureSystem</title>
  <link href='https://fonts.googleapis.com/css?family=Raleway:200,400,800' rel='stylesheet' type='text/css'>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css">
  <link rel='stylesheet' href='https://www.marcoguglie.it/Codepen/AnimatedHeaderBg/demo-1/css/demo.css'>
  <link rel="stylesheet" href="css/style.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

</head>
<body>
  <div id="large-header" class="large-header">
      <form action="" method="post">
      <?php echo $msg; ?>
        <br />
        <h2 class="main-title">Welcome to our supersecure "Ecorp Login System"Â©</h2>
        	<input class="username" type="text" placeholder="username" name="username"/>
        	<input class="password" type="password" placeholder="token" name="token"/>
        <button class="ok_button" type="submit" id="pwdButton" value="Login" name="SubmitButton">Login</button>
      </form>
    <a href="source.txt"> View Source </a>
    <canvas id="demo-canvas"></canvas>
  </div>
<script src='https://www.marcoguglie.it/Codepen/AnimatedHeaderBg/demo-1/js/EasePack.min.js'></script>
<script src='https://www.marcoguglie.it/Codepen/AnimatedHeaderBg/demo-1/js/rAF.js'></script>
<script src='https://www.marcoguglie.it/Codepen/AnimatedHeaderBg/demo-1/js/TweenLite.min.js'></script>
  <script src='js/index.js'></script>
</body>
</html>