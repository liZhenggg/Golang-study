<html>
<head>
<title>首页--输入个人信息</title>
</head>
<body>
<div>
	<h3>填写个人信息</h3>
	<form action="/userInfoEdit" method="POST">
		<table>
			<tr>
				<td width="100px">昵称*</td>
				<td width="300px"><input type="text" name="nickname" value="Zac" requied="true" /></td>
			</tr>
			
			<tr>
				<td width="100px">中文名*</td>
				<td width="300px"><input type="text" name="name_CN" value="大大" requied="true" /></td>
			</tr>
			
			<tr>
				<td width="100px">英文名*</td>
				<td width="300px"><input type="text" name="name_EN" value="dada" requied="true" /></td>
			</tr>
			
			<tr>
				<td width="100px">年龄*</td>
				<td width="300px"><input type="text" name="age" value="33" requied="true" /></td>
			</tr>
			
			<tr>
				<td width="100px">性别</td>
				<td width="300px">
					<label><input type="radio" name="gender" value="1" />男</label> 
					<label><input type="radio" name="gender" value="0" />女</label> 
					<label><input type="radio" name="gender" value="2" />保密</label> 
				</td>
			</tr>
			
			<tr>
				<td width="100px">身份证号</td>
				<td width="300px"><input type="text" name="idcard" /></td>
			</tr>
			
			<tr>
				<td width="100px">电子邮件</td>
				<td width="300px"><input type="text" name="email" /></td>
			</tr>
			
			<tr>
				<td width="100px">手机号码</td>
				<td width="300px"><input type="text" name="mobile" /></td>
			</tr>
	
			<tr>
				<td width="100px">国家</td>
				<td width="300px">
					<select name="nationality" >
					  <option value="0">其他</option>
					  <option value="1" >中国</option>
					  <option value="2">美国</option>
					  <option value="3">日本</option>
					  <option value="4">巴布亚新几内亚</option>
					</select>
				</td>
			</tr>
			
			<tr>
				<td width="100px">爱好</td>
				<td width="300px">
					<label><input type="checkbox" name="hobby" value="1" />旅游</label/>
					<label><input type="checkbox" name="hobby" value="2"/>美食</label/>
					<label><input type="checkbox" name="hobby" value="3"/>影视</label/>
					<label><input type="checkbox" name="hobby" value="4"/>运动</label/>
					<label><input type="checkbox" name="hobby" value="5"/>摄影</label/>
					<label><input type="checkbox" name="hobby" value="6"/>绘画</label/>
					<label><input type="checkbox" name="hobby" value="7"/>音乐</label/>
					<label><input type="checkbox" name="hobby" value="8"/>游戏</label/>
					<label><input type="checkbox" name="hobby" value="9"/>书籍</label/>
				</td>
			</tr>
			
			<tr>
				<td width="100px">生日</td>
				<td width="300px"><input type="text" name="birthday" /></td>
			</tr>
		</table>
		<input type="submit" value="提交" />
	</form>
</div>
</body>
</html>