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

func (i *Impl) Iter(starting []byte, limit int) ([]Pair, error) {
	var rdOpt = grocksdb.NewDefaultReadOptions()
	it := i.db.NewIterator(rdOpt)
	it.Seek(starting)
	res := make([]Pair, 0, limit)
    var idx = 0;
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
