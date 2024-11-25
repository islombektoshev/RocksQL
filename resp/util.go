package resp

import (
	"math/rand"

	"github.com/islombektoshev/RocksQL/check"
	"github.com/islombektoshev/RocksQL/util"
)

func RandomRESPCmd(rn *rand.Rand, maxArrLen, maxCmdLen int) RESPValue {
	cmd := RESPValue{Type: TypeArray}
	lenOfCmds := rn.Intn(maxArrLen) + 1
	for j := 0; j <= lenOfCmds; j++ {
		l := rn.Intn(maxCmdLen)
		cmdBytes := []byte{}
		for k := 0; k <= l; k++ {
			b := byte(rn.Intn(1 << 8))
			cmdBytes = append(cmdBytes, b)
		}
		cmd.Array = append(cmd.Array, RESPValue{
			Type:        TypeBulkString,
			StringValue: string(cmdBytes),
		})
	}
	check.That(cmd.Type == TypeArray)
	check.That(len(cmd.Array) > 0)
	return cmd
}

var respTypes = [5]byte{TypeSimpleString, TypeSimpleError, TypeInteger, TypeBulkString, TypeArray}

func RandomRESPValue(rn *rand.Rand, maxBound int, depth int) RESPValue {
	tayp := respTypes[rn.Intn(5)]
	if tayp == TypeArray && depth <= 0 {
		tayp = respTypes[rn.Intn(4)]
	}

	res := RESPValue{Type: tayp}

	switch tayp {
	case TypeSimpleString, TypeSimpleError:
		res.StringValue = util.RandString(rn, maxBound, []byte{'\r', '\n'})
		return res
	case TypeBulkString:
		res.StringValue = util.RandString(rn, maxBound, []byte{})
		return res
	case TypeInteger:
		res.IntValue = rn.Int()
		return res
	case TypeArray:
		count := rn.Intn(maxBound)
		res.Array = make([]RESPValue, count)
		for i := 0; i < count; i++ {
			res.Array[i] = RandomRESPValue(rn, maxBound, depth-1)
		}
	}
	check.That(res.Type != 0)
	switch tayp {
	case TypeSimpleError, TypeSimpleString, TypeBulkString:
		check.That(res.IntValue == 0)
		check.That(len(res.Array) == 0)
		break
	case TypeArray:
		check.That(res.IntValue == 0)
		check.That(len(res.StringValue) == 0)
		break
	case TypeInteger:
		check.That(len(res.Array) == 0)
		check.That(len(res.StringValue) == 0)
		break
	}
	return res
}
