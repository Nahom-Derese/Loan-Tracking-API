package emailutil

import "fmt"

func PasswordResetTemplate(url string) string {
	return fmt.Sprintf(
		`<html>
	<head>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				background-color: #e0f7fa;
				color: #004d40;
				margin: 0;
				padding: 0;
			}
			.container {
				width: 100%v;
				max-width: 600px;
				margin: 0 auto;
				background-color: #ffffff;
				padding: 25px;
				box-shadow: 0 0 15px rgba(0, 77, 64, 0.2);
				border-radius: 10px;
			}
			.header {
				text-align: center;
				padding: 15px 0;
				background-color: #00796b;
				color: #ffffff;
				border-radius: 10px 10px 0 0;
			}
			.content {
				padding: 25px;
				text-align: center;
			}
			.content p {
				font-size: 17px;
				line-height: 1.6;
			}
			.content a {
				display: inline-block;
				margin-top: 25px;
				padding: 12px 25px;
				color: #ffffff;
				background-color: #00796b;
				text-decoration: none;
				border-radius: 5px;
				font-weight: bold;
				transition: background-color 0.3s;
			}
			.content a:hover {
				background-color: #004d40;
			}
			.footer {
				text-align: center;
				padding: 15px 0;
				font-size: 13px;
				color: #004d40;
				border-top: 1px solid #b2dfdb;
				margin-top: 20px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Password Reset</h1>
			</div>
			<div class="content">
				<p>Please click the button below to reset your password:</p>
				<a href='%v'>Reset Password</a>
				<p>Thank you!</p>
			</div>
			<div class="footer">
				<p>&copy; 2024 Your Company. All rights reserved.</p>
			</div>
		</div>
	</body>
</html>`,
		"%",
		url,
	)
}
