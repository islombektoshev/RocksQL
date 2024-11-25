package engine

import "github.com/linxGnu/grocksdb"

type DBContext struct {
	DB          *grocksdb.TransactionDB
	DBOptions   *grocksdb.Options
	TXDBOptions *grocksdb.TransactionDBOptions
}

func (c *DBContext) Close() {
	c.DB.Close()
	c.DBOptions.Destroy()
	c.TXDBOptions.Destroy()
}

func OpenDB(path string) (DBContext, error) {
	var opt = grocksdb.NewDefaultOptions()
	opt.SetCreateIfMissing(true)
	var txdbOpt = grocksdb.NewDefaultTransactionDBOptions()
	var db, err = grocksdb.OpenTransactionDb(opt, txdbOpt, path)
	if err != nil {
		opt.Destroy()
		txdbOpt.Destroy()
		return DBContext{}, err
	}
	return DBContext{
		DB:          db,
		DBOptions:   opt,
		TXDBOptions: txdbOpt,
	}, nil
}
