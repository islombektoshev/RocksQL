package handles

import (
	"fmt"
	"strconv"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/resp"
)

type CmdCtx struct {
	Ctx    engine.SessionContext
	Engine engine.Engine
	Cmds   []string
}

func PUT(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) != 3 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format PUT. Example: `PUT 1 value`",
		}
	}
	var key = c.Cmds[1]
	var val = c.Cmds[2]
	var res = c.Engine.Put([]byte(key), []byte(val))
	if res.Err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: fmt.Errorf("Error occurding during PUT: %+v", res.Err).Error(),
		}
	} else {
		return resp.RESPValue{
			Type:        resp.TypeSimpleString,
			StringValue: "OK",
		}
	}
}

func GET(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) != 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format GET. Example: `GET 1`",
		}
	}
	var key = c.Cmds[1]
	var res = c.Engine.Get([]byte(key))
	if res.Err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: fmt.Errorf("Error occurding during GET: %+v", res.Err).Error(),
		}
	}
	return resp.RESPValue{
		Type:        resp.TypeBulkString,
		StringValue: string(res.Data),
	}
}

func ITER(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) > 3 || len(c.Cmds) < 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format ITER. Example: `ITER count [start]`",
		}
	}
	var limit, err = strconv.Atoi(c.Cmds[1])
	if err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "cannot parse count",
		}
	}
	var startingKey []byte = nil
	if len(c.Cmds) > 2 {
		startingKey = []byte(c.Cmds[2])
	}
	pairs, err := c.Engine.Iter(startingKey, limit)
	if err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: fmt.Errorf("Error occurding during GET: %+v", err).Error(),
		}
	}
	res := resp.RESPValue{
		Type: resp.TypeArray,
	}
	for _, p := range pairs {
		res.Array = append(res.Array, resp.RESPValue{Type: resp.TypeBulkString, StringValue: string(p.Key)})
		res.Array = append(res.Array, resp.RESPValue{Type: resp.TypeBulkString, StringValue: string(p.Val)})
	}
	return res

}

func DEL(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) != 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format DEL. Example: `DEL 1`",
		}
	}
	var key = c.Cmds[1]
	var err = c.Engine.Del([]byte(key))
	if err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: fmt.Sprintf("Error occurding during DEL: %+v", err),
		}
	}
	return resp.RESPValue{
		Type:        resp.TypeSimpleString,
		StringValue: "ok",
	}
}

func EXISTS(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) != 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format EXISTS. Example: `EXISTS 1`",
		}
	}
	var key = c.Cmds[1]
	var res = c.Engine.Get([]byte(key))
	if res.Err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: fmt.Errorf("Error occurding during EXISTS: %+v", res.Err).Error(),
		}
	}
	int := 0
	if res.Exists {
		int = 1
	}
	return resp.RESPValue{
		Type:     resp.TypeInteger,
		IntValue: int,
	}
}

func TX_START(c CmdCtx) resp.RESPValue {
	if len(c.Cmds) != 1 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "TxStart do not expect argument",
		}
	}
	index, err := c.Engine.TxStart(c.Ctx)
	if err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Tx Start Error: " + err.Error(),
		}
	}

	return resp.RESPValue{
		Type:        resp.TypeBulkString,
		StringValue: strconv.FormatUint(index, 10),
	}

}

func parseUint16(input string) (uint16, error) {
	val, err := strconv.ParseUint(input, 10, 16)
	return uint16(val), err
}

func parseUint64(input string) (uint64, error) {
	val, err := strconv.ParseUint(input, 10, 64)
	return uint64(val), err
}
