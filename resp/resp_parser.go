package resp

import (
	"fmt"
	"io"
	"strconv"
)

type FeedRes struct {
	Status int
	Cmds   []string
}

const (
	ResStatusContinue = iota
	ResStatusError
	ResStatusParsed
)

const (
	TypeSimpleString byte = '+'
	TypeSimpleError  byte = '-'
	TypeInteger      byte = ':'
	TypeBulkString   byte = '$'
	TypeArray        byte = '*'
)

type RESPParser struct {
	reader *RESPReader
	buffer []byte // Buffer for partial data
}

func NewParser(
	reader io.Reader,
) *RESPParser {
	return &RESPParser{
		reader: NewRESPReader(reader),
	}
}

type RESPValue struct {
	Type        byte
	StringValue string
	IntValue    int
	Array       []RESPValue
}

func (r *RESPParser) ReadNext() (RESPValue, error) {
	firstByte, err := r.reader.ReadByte()
	if err != nil {
		return RESPValue{}, err
	}

	var res RESPValue
	res.Type = firstByte
	switch firstByte {
	case TypeArray:
		lenAsStr, err := r.reader.ReadLine()
		if err != nil {
			return RESPValue{}, err
		}

		count, err := strconv.Atoi(lenAsStr)
		if err != nil {
			return RESPValue{}, err
		}
		if count < 0 {
			return RESPValue{}, fmt.Errorf("len of array is less then 0")
		}
		res.Array = make([]RESPValue, count)
		for i := 0; i < count; i++ {
			item, err := r.ReadNext()
			if err != nil {
				return RESPValue{}, err
			}
			res.Array[i] = item
		}
		break
	case TypeBulkString:
		lenAsStr, err := r.reader.ReadLine()
		if err != nil {
			return RESPValue{}, err
		}

		count, err := strconv.Atoi(lenAsStr)
		if err != nil {
			return RESPValue{}, err
		}
		if count < 0 {
			return RESPValue{}, fmt.Errorf("len of bluk string is less then 0")
		}
		stringBytes := make([]byte, count)
		n, err := r.reader.Read(stringBytes)
		if err != nil {
			return RESPValue{}, err
		}
		if n < count {
			return RESPValue{}, fmt.Errorf("did not as read as much as defined len of string")
		}

		line, err := r.reader.ReadLine()
		if err != nil {
			return RESPValue{}, err
		}
		if len(line) > 0 {
			return RESPValue{}, fmt.Errorf("recveign bulk string is not termetated with \\r\\n")
		}
		res.StringValue = string(stringBytes)
		break
	case TypeInteger:
		lineAsString, err := r.reader.ReadLine()
		if err != nil {
			return RESPValue{}, err
		}
		intVal, err := strconv.Atoi(lineAsString)
		if err != nil {
			return RESPValue{}, err
		}
		res.IntValue = intVal
		break

	case TypeSimpleString, TypeSimpleError:
		line, err := r.reader.ReadLine()
		fmt.Printf("reading line for %v, line: %s\n", rune(firstByte), line)
		if err != nil {
			return RESPValue{}, err
		}
		res.StringValue = line
		break
	default:
		return RESPValue{}, fmt.Errorf("unexpected data type prefix")
	}
	return res, nil

}
