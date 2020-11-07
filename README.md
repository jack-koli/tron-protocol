# tron/protocol

## 原因和目的

明显 tron 团队不太注重 golang ,应该也没有测试过，我这边用到 tron protocol,因此自己修改和编译一份


## 修改

每个 proto 文件的 option go_package 部分

## 使用

参考 https://github.com/jack-koli/tron-demo


```go
// main.go
package main

import (
	"context"
	"github.com/jack-koli/tron-protocol/api"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	//grpc.shasta.trongrid.io:50051
	addr := "grpc.shasta.trongrid.io:50051"
	conn, err := grpc.Dial( addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to %s, err:%v" , addr, err)
	}
	client := api.NewWalletClient(conn)

	block, err := walletClient.GetNowBlock(context.Background(), new(EmptyMessage))
	if err != nil {
		t.Errorf("get block err: %v", err)
	}
	rawData := block.BlockHeader.GetRawData()


	//timestamp := time.Time

	log.Printf("blockNumber : %d, timestamp : %d", rawData.Number, rawData.Timestamp)
}

```