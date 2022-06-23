###微信消息推送

服务：pushplus；官网：   http://www.pushplus.plus

下载
````
go get github.com/ANMP0042/wxPush
````

获取token
````
https://pushplus.hxtrip.com 注册
````

使用
````go

// 实例化
p,err := wxPush.NewPush(token)
if err != nil {
    fmt.Println("err = ", err)
    return
}

// 发送消息
if err = p.Send(Args{Title: "这是title", Content: "这里是contnet"}); err != nil {
    fmt.Println("推送错误 = ", err)
} else {
	fmt.Println("推送成功")
}

````


@更新日志：http://pushplus.hxtrip.com服务停了，更换成http://www.pushplus.plus  2022-06-23