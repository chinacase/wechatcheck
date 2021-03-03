# wechatcheck

  微信域名检测
------
# install
```
go get -u github.com/chinacase/wechatcheck
```

# example
------
```
package main

import (
	"github.com/chinacase/wechatcheck"
	"fmt"
)



func main() {

	func main() {

	fmt.Println("hello world")
	r := wechat.CheckURL("www.baidu.com/")
	//r := wechat.CheckURL("https://www.baidu.com/") //或者
	fmt.Println(r)
}
}
```
```
//Result 返回信息
type Result struct {
	Status int    //状态 1 通过 0 不通过 2 未知
	Msg    string //提示
}

```
