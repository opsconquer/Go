# Go
学习代码总结
Go语言里面提供了一个完善的net/http包，通过
http包可以很方便的就搭建起来一个可以运行的Web服务。同时使用这个包能很简单地对Web的路由，静
态文件，模版，cookie等数据进行设置和操作。
package main

import (
  "fmt" 
  "net/http"
  "strings"
  "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm() //解析参数，默认是不会解析的
  fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key:", k) 
    fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
http.HandleFunc("/", sayhelloName) //设置访问的路由
err := http.ListenAndServe(":9090", nil) //设置监听的端口
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

在浏览器输入http://localhost:9090

可以看到浏览器页面输出了Hello astaxie!

可以换一个地址试试： http://localhost:9090/?url_long=111&url_long=222
