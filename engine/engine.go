package engine

type OpRes struct {
	Err error
}

type DataRes struct {
	OpRes
	Data   []byte
	Exists bool
}

type Pair struct {
	Key []byte
	Val []byte
}
type Engine interface {
	Put([]byte, []byte) OpRes
	Get([]byte) DataRes
	Del([]byte) error
	TxStart(SessionContext) (uint64, error)
	Iter(starting []byte, limit int) ([]Pair, error)
}
