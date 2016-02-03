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







