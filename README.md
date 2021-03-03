# wechatcheck

  微信相关的方法
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

	fmt.Println("hello world")
	r := wechat.CheckURL("www.52bd.net/code/256.html")
	fmt.Println(r)
}
```
```
//Result 返回信息
type Result struct {
	Status int    //状态 1 通过 0 不通过 2 未知
	Msg    string //提示
}

```
