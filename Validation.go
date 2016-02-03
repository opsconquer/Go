必填字段：
if len(r.Form["username"][0])==0{
//为空的处理
}

数字字段：
getint,err:=strconv.Atoi(r.Form.Get("age"))
if err!=nil{
//数字转化出错了，那么可能就不是数字
}
//接下来就可以判断这个数字的大小范围了
if getint >100 {
//太大了
}
正则方式
if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
return false
}

中文字段：
if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
return false
}
对于中文我们目前有两种方式来验证，可以使用 unicode 包提供
的 func Is(rangeTab *RangeTable, r rune) bool 来验证。

英文字段
if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
return false
}

电子邮件地址
if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
fmt.Println("no")
}else{
fmt.Println("yes")
}

手机号码
if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
return false
}

下来菜单
<select name="fruit">
<option value="apple">apple</option>
<option value="pear">pear</option>
<option value="banane">banane</option>
</select>
验证
slice:=[]string{"apple","pear","banane"}
for _, v := range slice {
if v == r.Form.Get("fruit") {
return true
}
}
return false

单选按钮
<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女
验证
slice:=[]int{1,2}
for _, v := range slice {
if v == r.Form.Get("gender") {
return true
}
}
return false

复选框
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
验证:
slice:=[]string{"football","basketball","tennis"}
a:=Slice_diff(r.Form["interest"],slice)
if a == nil{
return true
}
return false

日期和时间
t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
fmt.Printf("Go launched at %s\n", t.Local())

身份证号码
//验证15位身份证，15位的是全部数字
if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
return false
}
//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
return false
}


fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

import "text/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
输出
Hello, <script>alert('you have been pwned')</script>!

import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", template.HTML("<script>alert('you have been pwned')</script>"))
输出
Hello, <script>alert('you have been pwned')</script>!















