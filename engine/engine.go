package engine

type OpRes struct {
	Err error
}

type DataRes struct {
	OpRes
	Data []byte
}

type Engine interface {
	Put([]byte, []byte) OpRes
	Get([]byte) DataRes
}
