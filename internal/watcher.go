package internal

import (
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"time"
)

const interval_check_duration = 10 * time.Second // Maximum

type TxCallbackFn func()

type TxCallback struct {
	EmittedBlock int // When did the transaction was emitted in block height
	Tx           *types.Transaction
	Fn           TxCallbackFn
}

func CheckModules(ctx *Context) error {
	modules, err := ctx.RpcConn.SupportedModules()
	if err != nil {
		return err
	}
	log.Printf("Supported modules: %q", modules)
	return nil
}

func Watcher(ctx *Context) {
	err := CheckModules(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(interval_check_duration)
	var txs []TxCallback
	for {
		select {
		case <-ticker.C:
			//TODO cleanup
		case element := <-ctx.WatchChan:
			txs = append(txs, element)
		}
	}
}
