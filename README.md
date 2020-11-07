# tron/protocol

## 原因和目的

明显 tron 团队不太注重 golang ,应该也没有测试过，我这边用到 tron protocol,因此自己修改和编译一份


## 修改

每个 proto 文件的 option go_package 部分

## 使用


```go
import (
	"github.com/jack-koli/tron-protocol/api"
)
```