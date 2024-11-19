package server_test

import (
	"testing"

	"github.com/islombektoshev/RocksQL/resp"
	"github.com/islombektoshev/RocksQL/server"
	"github.com/stretchr/testify/assert"
)


func TestRESPValueToCmd(t *testing.T) {
    cmds, err := server.RespValueToCmd(resp.RESPValue{
        Type: resp.TypeArray,
        Array: []resp.RESPValue{
            {
                Type: resp.TypeBulkString,
                StringValue: "GET",
            },
            {
                Type: resp.TypeBulkString,
                StringValue: "a",
            },
        },
    })
    assert.NoError(t, err)
    assert.Equal(t, []string{"GET", "a"}, cmds)
}
