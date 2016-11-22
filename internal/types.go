package internal

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jinzhu/gorm"
)

type Context struct {
	Db        *gorm.DB
	RpcConn   *rpc.Client
	WatchChan chan TxCallback
	OpCount   int64
}
