package resp_test

import (
	"testing"

	"github.com/islombektoshev/RocksQL/resp"
	"github.com/stretchr/testify/assert"
)

func TestSerilizer(t *testing.T) {
	data := []byte("*4\r\n$3\r\nGET\r\n:123\r\n+hello\r\n-error-content\r\n")
	val, err := resp.SerilizeRespValue(
		resp.RESPValue{
			Type: resp.TypeArray,
			Array: []resp.RESPValue{
				{
					Type:        resp.TypeBulkString,
					StringValue: "GET",
				},
				{
					Type:     resp.TypeInteger,
					IntValue: 123,
				},
				{
					Type:        resp.TypeSimpleString,
					StringValue: "hello",
				},
				{
					Type:        resp.TypeSimpleError,
					StringValue: "error-content",
				},
			},
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, data, val)

}
