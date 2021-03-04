# wechatcheck

  微信域名检测

# install
```
go get -u github.com/chinacase/wechatcheck
```

# example

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
+ Status 状态 1：安全性未知 2：危险网站 3：安全网站
