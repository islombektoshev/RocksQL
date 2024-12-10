package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/resp"
)

type Server struct {
	address  string
	handler  Handler
	listener net.Listener
}

//go:generate mockery --with-expecter --name=Handler --output=mocks --outpkg=mocks
type Handler interface {
	ExecuteCmd(ctx engine.SessionContext, cmds []string) resp.RESPValue
}

func NewServer(handler Handler, address string) *Server {
	return &Server{
		handler: handler,
		address: address,
	}
}

func (s *Server) StartListening() error {
	if s.listener != nil {
		return fmt.Errorf("Already listening")
	}

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.listener = listener
	return nil
}

func (s *Server) Addr() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}
	return ""
}
func (s *Server) handeConn(ctx engine.SessionContext, conn net.Conn) {
	reader := resp.NewRESPReader(conn)
	parser := resp.NewParser(reader)

readBlk:
	for {
		var cmds []string
		var result resp.RESPValue
		respValue, err := parser.ReadNext()
		if err != nil {
			if err == io.EOF {
				break readBlk
			} else {
				result = resp.RESPValue{
					Type:        resp.TypeSimpleError,
					StringValue: err.Error(),
				}
				goto respondeWithResult
			}
		}

		cmds, err = RespValueToCmd(respValue)
		if err != nil {
			result = resp.RESPValue{
				Type:        resp.TypeSimpleError,
				StringValue: err.Error(),
			}

			goto respondeWithResult
		}
		result = s.handler.ExecuteCmd(ctx, cmds)

	respondeWithResult:
		resByts, err := resp.SerializeRespValue(result)
		if err != nil {
			log.Fatal("[BUG] Cannot parse simple RESPValue")
		}
		_, err = conn.Write(resByts)
		if err != nil {
			slog.Warn("Closing connection, Connection returned error.", "error", err)
			conn.Close()
			break readBlk
		}

	}
}

func (s *Server) AcceptContinously(ctx context.Context, MAX_CONN int) {
	var semaphore = make(chan int, MAX_CONN)
	var byebye = false
	go func() {
		<-ctx.Done()
		byebye = true
		slog.Info("Shutting Down Sever")
		err := s.listener.Close()
		if err != nil {
			slog.Error("Unable to close server", "error", err)
		}
	}()

	for {
		semaphore <- 0

		conn, err := s.listener.Accept()
		if err != nil {
			if byebye {
				slog.Info("Server: connection closed: err {}", "error", err)
				return
			} else {
				slog.Error("Server: cannot accpet connection: err {}", "error", err)
			}
		}

		slog.Debug("New Connection created", "raddr", conn.RemoteAddr())

		go func() {
			defer func() {
				<-semaphore
			}()

			s.handeConn(engine.SessionContext{
				Context: ctx,
				Vars:    engine.NewSession(),
			}, conn)
		}()
	}
}

func RespValueToCmd(val resp.RESPValue) ([]string, error) {
	if val.Type != resp.TypeArray {
		return nil, fmt.Errorf("CMD should be wapped with array")
	}
	res := make([]string, 0)
	for _, item := range val.Array {
		if item.Type != resp.TypeBulkString {
			return nil, fmt.Errorf("CMD shouble bulk string")
		}
		res = append(res, item.StringValue)
	}
	return res, nil
}
