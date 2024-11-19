package resp

import (
	"fmt"
	"strconv"
)

func SerilizeRespValue(val RESPValue) ([]byte, error) {
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
		for _, c := range val.StringValue {
			if c == '\r' {
				res = append(res, '\\', 'r')
			} else if c == '\n' {
				res = append(res, '\\', 'n')
			} else {
				res = append(res, byte(c))
			}
		}
		res = append(res, '\r', '\n')
		return res, nil
	case TypeInteger:
		res = append(res, []byte(strconv.FormatInt(int64(val.IntValue), 10))...)
		res = append(res, '\r', '\n')
		return res, nil
	case TypeBulkString:
		res = append(res, []byte(strconv.FormatInt(int64(len(val.StringValue)), 10))...)
		res = append(res, '\r', '\n')
		res = append(res, []byte(val.StringValue)...)
		res = append(res, '\r', '\n')
	case TypeArray:
		res = append(res, []byte(strconv.FormatInt(int64(len(val.Array)), 10))...)
		res = append(res, '\r', '\n')
		for _, item := range val.Array {
			itemVal, err := SerilizeRespValue(item)
			if err != nil {
				return nil, err
			}
			res = append(res, itemVal...)
		}
	}
	return res, nil
}
