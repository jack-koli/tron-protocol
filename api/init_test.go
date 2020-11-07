package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	log.Println("TestMain")
	os.Exit(m.Run())
}


func TestNewWalletClient(t *testing.T) {
	conn, err := grpc.Dial("grpc.shasta.trongrid.io:50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("can not connect ot grpc node, err:%v", err)
	}

	walletClient := NewWalletClient(conn)
	block, err := walletClient.GetNowBlock(context.Background(), new(EmptyMessage))
	if err != nil {
		t.Errorf("get block err: %v", err)
	}
	rawData := block.BlockHeader.GetRawData()


	//timestamp := time.Time

	log.Printf("blockNumber : %d, timestamp : %d", rawData.Number, rawData.Timestamp)
}
