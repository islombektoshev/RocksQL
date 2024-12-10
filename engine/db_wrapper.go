package engine

import (
	"context"
	"sync/atomic"

	"github.com/linxGnu/grocksdb"
)

type EngineContext struct {
	DB          *grocksdb.TransactionDB
	DBOptions   *grocksdb.Options
	TXDBOptions *grocksdb.TransactionDBOptions
	TxList      map[uint64]TxHolder
	_counter    *atomic.Uint64
}

type TxHolder struct {
	Tx    *grocksdb.Transaction
	Index uint64
}

type SessionContext struct {
	context.Context
	Vars   *SessionVars
	TxList map[uint64]TxHolder
}

type SessionVars struct {
	DB      string
	ColFam  string
	ParenTx uint64
}

const SessionKey = "engine-session"

func NewSession() *SessionVars {
	return &SessionVars{}
}
func (c *EngineContext) Close() {
	c.DB.Close()
	c.DBOptions.Destroy()
	c.TXDBOptions.Destroy()
}

func (c *EngineContext) NextInt() uint64 {
	return c._counter.Add(1)
}
func OpenDB(path string) (EngineContext, error) {
	var opt = grocksdb.NewDefaultOptions()
	opt.SetCreateIfMissing(true)
	opt.EnableStatistics()
	var txdbOpt = grocksdb.NewDefaultTransactionDBOptions()
	var db, err = grocksdb.OpenTransactionDb(opt, txdbOpt, path)
	if err != nil {
		opt.Destroy()
		txdbOpt.Destroy()
		return EngineContext{}, err
	}
	ctx := EngineContext{
		DB:          db,
		DBOptions:   opt,
		TXDBOptions: txdbOpt,
		_counter:    &atomic.Uint64{},
	}
	ctx.NextInt()
	return ctx, nil
}
