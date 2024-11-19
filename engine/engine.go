package engine

type OpRes struct {
	Err error
}

type DataRes struct {
	OpRes
	Data []byte
}

type Pair struct {
	Key []byte
	Val []byte
}
type Engine interface {
	Put([]byte, []byte) OpRes
	Get([]byte) DataRes
	Iter(starting []byte, limit int) ([]Pair, error)
}
