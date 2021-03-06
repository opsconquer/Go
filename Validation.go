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

转义的例子
import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
转义之后的输出：
Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!


防止多次递交表单
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
用户名:<input type="text" name="username">
密码:<input type="password" name="password">
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="登陆">
增加了一个隐藏字段token ，这个值我们通过MD5(时间戳)来获取惟一值
func login(w http.ResponseWriter, r *http.Request) {
fmt.Println("method:", r.Method) //获取请求的方法
if r.Method == "GET" {
crutime := time.Now().Unix()
h := md5.New()
io.WriteString(h, strconv.FormatInt(crutime, 10))
token := fmt.Sprintf("%x", h.Sum(nil))
t, _ := template.ParseFiles("login.gtpl")
t.Execute(w, token)
} else {
//请求的是登陆数据，那么执行登陆的逻辑判断
r.ParseForm()
token := r.Form.Get("token")
if token != "" {
//验证token的合法性
} else {
//不存在token报错
}
fmt.Println("username length:", len(r.Form["username"][0]))
fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
}
}

处理文件上传
要使表单能够上传文件，首先第一步就是要添加form的enctype 属性， enctype 属性有如下三种情况:
application/x-www-form-urlencoded 表示在发送前编码所有字符（默认）
multipart/form-data 不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
text/plain 空格转换为 "+" 加号，但不对特殊字符编码。
http.HandleFunc("/upload", upload)
// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
fmt.Println("method:", r.Method) //获取请求的方法
if r.Method == "GET" {
crutime := time.Now().Unix()
h := md5.New()
io.WriteString(h, strconv.FormatInt(crutime, 10))
token := fmt.Sprintf("%x", h.Sum(nil))
t, _ := template.ParseFiles("upload.gtpl")
t.Execute(w, token)
} else {
r.ParseMultipartForm(32 << 20)
file, handler, err := r.FormFile("uploadfile")
if err != nil {
fmt.Println(err)
return
}
defer file.Close()
fmt.Fprintf(w, "%v", handler.Header)
f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
fmt.Println(err)
return
}
defer f.Close()
io.Copy(f, file)
}
}


客户端上传文件
package main
import (
"bytes"
"fmt"
"io"
"io/ioutil"
"mime/multipart"
"net/http"
"os"
)
func postFile(filename string, targetUrl string) error {
bodyBuf := &bytes.Buffer{}
bodyWriter := multipart.NewWriter(bodyBuf)
//关键的一步操作
fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
if err != nil {
fmt.Println("error writing to buffer")
return err
}
//打开文件句柄操作
fh, err := os.Open(filename)
if err != nil {
fmt.Println("error opening file")
return err
}
defer fh.Close()
//iocopy
_, err = io.Copy(fileWriter, fh)
if err != nil {
return err
}
contentType := bodyWriter.FormDataContentType()
bodyWriter.Close()
resp, err := http.Post(targetUrl, contentType, bodyBuf)
if err != nil {
return err
}
defer resp.Body.Close()
resp_body, err := ioutil.ReadAll(resp.Body)
if err != nil {
return err
}
fmt.Println(resp.Status)
fmt.Println(string(resp_body))
return nil
}
// sample usage
func main() {
target_url := "http://localhost:9090/upload"
filename := "./astaxie.pdf"
postFile(filename, target_url)














