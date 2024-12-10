package server_test

import (
	"context"
	"log/slog"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/islombektoshev/RocksQL/check"
	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/resp"
	"github.com/islombektoshev/RocksQL/server"
	"github.com/islombektoshev/RocksQL/server/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRESPValueToCmd(t *testing.T) {
	cmds, err := server.RespValueToCmd(resp.RESPValue{
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
	})
	assert.NoError(t, err)
	assert.Equal(t, []string{"GET", "a"}, cmds)
}

func TestSeverTest_CanStart(t *testing.T) {
	handler := mocks.NewHandler(t)
	svr := server.NewServer(handler, "127.0.0.1:")
	ctx, cancel := context.WithCancel(context.Background())
	svr.StartListening()
	go func() {
		svr.AcceptContinously(ctx, 10)
	}()
	cancel()
}

func TestSeverTest_RandomCommands(t *testing.T) {
	slog.Info("Starting Sever Random Test")
	handler := mocks.NewHandler(t)

	svr := server.NewServer(handler, "127.0.0.1:")
	ctx, cancel := context.WithCancel(context.Background())
	svr.StartListening()
	var cancelDeferWorked = false
	defer func() {
		cancelDeferWorked = true
		slog.Info("Canceling Server")
		cancel()
	}()

	go func() {
		slog.Info("Accept Continously")
		svr.AcceptContinously(ctx, 10)
		check.That(cancelDeferWorked)
	}()

	rnSource := rand.NewSource(1732157844)
	rn := rand.New(rnSource)

	conn, err := net.Dial("tcp", svr.Addr())
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()
	require.NoError(t, err)

	var reader = resp.NewRESPReader(conn)
	var parser = resp.NewParser(reader)

	var lastCmd *resp.RESPValue
	var lastResp resp.RESPValue
	var responseId = 0
	handler.EXPECT().ExecuteCmd(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx engine.SessionContext, cmds []string) resp.RESPValue {
			slog.Info("Server: Receiving command")
			require.Equal(t, len(lastCmd.Array), len(cmds))
			for i, cmd := range cmds {
				require.Equal(t, resp.TypeBulkString, lastCmd.Array[i].Type)
				require.Equal(t, lastCmd.Array[i].StringValue, cmd)
			}
			lastResp = resp.RandomRESPValue(rn, 100, 2)
			check.That(lastResp.Type != 0)
			slog.Info("Server: Sending response")
			responseId++
			return lastResp
		})

	for i := 0; i < 1000; i++ {
		check.That(!cancelDeferWorked)
		cmd := resp.RandomRESPCmd(rn, 100, 200)
		lastCmd = &cmd
		data, err := resp.SerializeRespValue(cmd)
		require.NoError(t, err)
		slog.Info("Client: Sending Command", "index", i)
		conn.Write(data)
		value, err := parser.ReadNext()
		slog.Info("Client: Receiving Response", "index", i)
		require.NoError(t, err)
		require.Equal(t, i, responseId-1, "Response id neet to be same")
		require.Equal(t, lastResp, value)
	}
}
