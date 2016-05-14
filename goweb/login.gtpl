<html>
<head>
	<title>login</title>
</head>
<body>
<input type="checkbox" name="interest" value="football">football
<input type="checkbox" name="interest" value="basketball">basketball
<form>
</form action="login" method="post">
	username:<input type="text" name="username">
	password:<input type="password" name="password">
	<input type="hidden" name="token" value="{{.}}" >
	<input type="submit" value="login">
</body>
</html>