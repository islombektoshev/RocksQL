package engine

import (
	"fmt"
	"github.com/linxGnu/grocksdb"
)

type Impl struct {
	EngineContext
}

func NewEngine(db EngineContext) Engine {
	return &Impl{EngineContext: db}
}

func (i *Impl) Put(key []byte, val []byte) OpRes {
	var wrOpt = grocksdb.NewDefaultWriteOptions()
	defer wrOpt.Destroy()

	err := i.DB.Put(wrOpt, key, val)

	if err != nil {
		return OpRes{Err: err}
	}
	return OpRes{Err: nil}
}

func (i *Impl) Get(key []byte) DataRes {
	var rdOpt = grocksdb.NewDefaultReadOptions()
	slice, err := i.DB.Get(rdOpt, key)
	if err != nil {
		return DataRes{OpRes{Err: err}, nil, false}
	}
	defer slice.Free()
	var bytes = slice.Data()
	var res = make([]byte, len(bytes))
	copy(res, bytes)
	return DataRes{OpRes{Err: nil}, res, slice.Exists()}
}

func (i *Impl) Iter(starting []byte, limit int) ([]Pair, error) {
	var rdOpt = grocksdb.NewDefaultReadOptions()
	it := i.DB.NewIterator(rdOpt)
	it.Seek(starting)
	res := make([]Pair, 0, limit)
	var idx = 0
	for it.Valid() && idx < limit {
		keyData := it.Key().Data()
		defer it.Key().Free()

		valData := it.Value().Data()
		defer it.Value().Free()

		var key = make([]byte, len(keyData))
		copy(key, keyData)

		var val = make([]byte, len(valData))
		copy(val, valData)
		res = append(res, Pair{
			Key: key,
			Val: val,
		})
		it.Next()
		idx++
	}
	return res, nil
}

func (i *Impl) Del(key []byte) error {
	wrOpt := grocksdb.NewDefaultWriteOptions()
	defer wrOpt.Destroy()
	return i.DB.Delete(wrOpt, key)
}

func (i *Impl) TxStart(ctx SessionContext) (uint64, error) {
	var parenTx *grocksdb.Transaction
	if ctx.Vars.ParenTx > 0 {
		parenTx = ctx.TxList[ctx.Vars.ParenTx].Tx
		if parenTx == nil {
			return 0, fmt.Errorf("Parent Tx not found")
		}
	}
	wrtOpt := grocksdb.NewDefaultWriteOptions()
	defer wrtOpt.Destroy()
	txOpt := grocksdb.NewDefaultTransactionOptions()
	defer txOpt.Destroy()

	tx := i.DB.TransactionBegin(wrtOpt, txOpt, parenTx)
	index := i.NextInt()
	txHoldr := TxHolder{
		Tx:    tx,
		Index: index,
	}

	ctx.TxList[index] = txHoldr
	i.TxList[index] = txHoldr
	return index, nil
}
