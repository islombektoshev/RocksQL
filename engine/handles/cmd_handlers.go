package handles

import (
	"strconv"

	"fmt"
	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/resp"
)

func PUT(eng engine.Engine, cmds []string) resp.RESPValue {
	if len(cmds) != 3 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format PUT. Example: `PUT 1 value`",
		}
	}
	var key = cmds[1]
	var val = cmds[2]
	var res = eng.Put([]byte(key), []byte(val))
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

func GET(eng engine.Engine, cmds []string) resp.RESPValue {
	if len(cmds) != 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format GET. Example: `GET 1`",
		}
	}
	var key = cmds[1]
	var res = eng.Get([]byte(key))
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

func ITER(eng engine.Engine, cmds []string) resp.RESPValue {
	if len(cmds) > 3 || len(cmds) < 2 {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "Wrong format ITER. Example: `ITER count [start]`",
		}
	}
	var limit, err = strconv.Atoi(cmds[1])
	if err != nil {
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "cannot parse count",
		}
	}
	var startingKey []byte = nil
	if len(cmds) > 2 {
		startingKey = []byte(cmds[2])
	}
	pairs, err := eng.Iter(startingKey, limit)
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

func parseUint16(input string) (uint16, error) {
	val, err := strconv.ParseUint(input, 10, 16)
	return uint16(val), err
}

func parseUint64(input string) (uint64, error) {
	val, err := strconv.ParseUint(input, 10, 64)
	return uint64(val), err
}
