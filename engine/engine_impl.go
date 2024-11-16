package engine

import "github.com/linxGnu/grocksdb"

type Impl struct {
	db *grocksdb.DB
}

func NewEngine(db *grocksdb.DB) Engine {
	return &Impl{db: db}
}

func (i *Impl) Put(key []byte, val []byte) OpRes {
	var wrOpt = grocksdb.NewDefaultWriteOptions()
	defer wrOpt.Destroy()

	err := i.db.Put(wrOpt, key, val)

	if err != nil {
		return OpRes{Err: err}
	}
	return OpRes{Err: nil}
}

func (i *Impl) Get(key []byte) DataRes {
	var rdOpt = grocksdb.NewDefaultReadOptions()
	slice, err := i.db.Get(rdOpt, key)
	if err != nil {
		return DataRes{OpRes{Err: err}, nil}
	}
	defer slice.Free()
	var bytes = slice.Data()
	var res = make([]byte, len(bytes))
	copy(res, bytes)
	return DataRes{OpRes{Err: nil}, res}
}
