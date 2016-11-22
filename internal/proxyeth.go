package internal

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func InitContext(wsURI string, pkey string, contractAddr string) (*Context, error) {
	var ctx Context
	ctx.WatchChan = make(chan TxCallback)
	conn, err := rpc.Dial(wsURI)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ctx.RpcConn = conn
	keyECDSA, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Fatalf("Failed to parse key: %v", err)
	}
	/*auth :=*/ bind.NewKeyedTransactor(keyECDSA)

	log.Println("Not used: ", contractAddr)

	return &ctx, nil
}
