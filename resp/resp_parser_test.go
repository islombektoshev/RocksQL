package resp_test

import (
	"testing"

	"io"

	"github.com/islombektoshev/RocksQL/resp"
	"github.com/stretchr/testify/assert"
)

func TestParserTestTrival_OK(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	parser := resp.NewParser(pipeReader, nil)
	go func() {
		val, err := parser.ReadNext()
		assert.NoError(t, err)
		assert.Equal(t, resp.RESPValue{
			Type: resp.TypeArray,
			Array: []resp.RESPValue{
				{
					Type:        resp.TypeBulkString,
					StringValue: "GET",
				},
				{
					Type:        resp.TypeBulkString,
					StringValue: "a",
				},
			},
		}, val)
	}()

	pipeWriter.Write([]byte("*2\r\n$3\r\nGET\r\n$1\r\na\r\n"))
}

func TestParserTestTrival_OK_2(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	parser := resp.NewParser(pipeReader, nil)
	go func() {
		pipeWriter.Write([]byte("*4\r\n$3\r\nGET\r\n:123\r\n+hello\r\n-error-content\r\n"))
	}()
	val, err := parser.ReadNext()
	assert.NoError(t, err)
	assert.Equal(t, resp.RESPValue{
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
	}, val)

}
