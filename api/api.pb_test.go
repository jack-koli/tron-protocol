package api

import (
	"context"
	"github.com/jack-koli/tron-protocol/core/contract"
	"github.com/jack-koli/tron-protocol/core"
	"github.com/jack-koli/tron-protocol/common/crypto/base58"
	"log"
	"testing"
)

func TestAccountResourceMessage_GetEnergyUsed(t *testing.T) {
	//walletClient := GetNewWalletClient()
	//walletClient
}

func TestGetAccount(t *testing.T)  {
	walletClient := GetNewWalletClient()
	acc := new(core.Account)
	addr := "TZ5dPxnxd4rRZb7nudcorifD9zfxi2NSRY"
	acc.Address = base58.Decode(addr)
	result,err := walletClient.GetAccount(context.Background(), acc)
	if err != nil {
		log.Fatalf("can not get result of %s", addr)
	}

	contract := new(contract.)
	log.Println("result:", result)
}
