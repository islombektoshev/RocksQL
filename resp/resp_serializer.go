package resp

import (
	"fmt"
	"strconv"

	"github.com/islombektoshev/RocksQL/check"
)

func SerializeRespValue(val RESPValue) ([]byte, error) {
	if val.Type != TypeArray &&
		val.Type != TypeSimpleString &&
		val.Type != TypeBulkString &&
		val.Type != TypeSimpleError &&
		val.Type != TypeInteger {
		return nil, fmt.Errorf("cannot serilize: unkown type: %s, byte: %d",
			string(val.Type), val.Type)
	}
	res := make([]byte, 1)
	res[0] = val.Type
	switch val.Type {
	case TypeSimpleString, TypeSimpleError:
		check.That(len(val.Array) == 0)
		check.That(val.IntValue == 0)
		check.That(len(res) == 1)
		replaces := 0
		var i = 0
		data := []byte(val.StringValue)
		check.That(len(val.StringValue) == len(data))
		for _, c := range data {
			if c == '\r' {
				return nil, fmt.Errorf("Simple String cannot contain \\r")
			} else if c == '\n' {
				return nil, fmt.Errorf("Simple String cannot contain \\n")
			} else {
				res = append(res, byte(c))
			}
			i++
		}
		check.That(len(res)-replaces-1 == len(val.StringValue))
		check.That(len(val.StringValue) == i)
		res = append(res, '\r', '\n')
		return res, nil
	case TypeInteger:
		check.That(len(val.StringValue) == 0)
		check.That(len(val.Array) == 0)
		res = append(res, []byte(strconv.FormatInt(int64(val.IntValue), 10))...)
		res = append(res, '\r', '\n')
		return res, nil
	case TypeBulkString:
		check.That(len(val.Array) == 0)
		check.That(val.IntValue == 0)
		id := strconv.FormatInt(int64(len(val.StringValue)), 10)
		res = append(res, []byte(id)...)
		res = append(res, '\r', '\n')
		res = append(res, []byte(val.StringValue)...)
		res = append(res, '\r', '\n')
		check.That(len(res) == len(id)+len(val.StringValue)+5)
	case TypeArray:
		check.That(len(val.StringValue) == 0)
		check.That(val.IntValue == 0)

		res = append(res, []byte(strconv.FormatInt(int64(len(val.Array)), 10))...)
		res = append(res, '\r', '\n')
		for _, item := range val.Array {
			itemVal, err := SerializeRespValue(item)
			if err != nil {
				return nil, err
			}
			res = append(res, itemVal...)
		}
	}
	return res, nil
}
