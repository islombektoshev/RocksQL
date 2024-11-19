package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"strings"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/resp"
)

type Server struct {
	eng      engine.Engine
	address  string
	listener net.Listener
}

func NewServer(eng engine.Engine, address string) *Server {
	return &Server{
		eng:     eng,
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

func (s *Server) AcceptContinously(cont context.Context) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			slog.Error("Server: cannot accpet connectition: err {}", err)
		}
		go func() {
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
				result = ExecuteCmd(cmds, s.eng)

			respondeWithResult:
				resByts, err := resp.SerilizeRespValue(result)
				if err != nil {
					log.Fatal("[BUG] Cannot parse simple RESPValue")
				}
				_, err = conn.Write(resByts)
				if err != nil {
					slog.Warn("Connection returned error: {}", err)
					slog.Warn("Closing connection")
					conn.Close()
					break readBlk
				}

			}
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

func ExecuteCmd(cmds []string, eng engine.Engine) resp.RESPValue {
	var firstToken = strings.ToUpper(cmds[0])
	switch firstToken {
	case "GET":
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
	case "PUT", "SET":
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
	default:
		return resp.RESPValue{
			Type:        resp.TypeSimpleError,
			StringValue: "unkown command: available commands: GET, PUT, SET",
		}
	}
}
