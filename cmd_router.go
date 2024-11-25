package main

import (
	"strings"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/engine/handles"
	"github.com/islombektoshev/RocksQL/resp"
)

type CmdRouter struct {
	Engine engine.Engine
}
var handlers = map[string]func(engine.Engine, []string) resp.RESPValue {
    "GET": handles.GET,
    "PUT": handles.PUT,
    "SET": handles.PUT,
    "ITER": handles.ITER,
}
var errorResponse resp.RESPValue;
func init () {
    var msg string = "unkown command was sent. available commands: "
    for key, _ := range handlers {
       msg = msg +  key + ", "
    }
    errorResponse = resp.RESPValue{
        Type: resp.TypeSimpleError,
        StringValue: msg,
    }
}

func (h *CmdRouter) ExecuteCmd(cmds []string) resp.RESPValue {
	var firstToken = strings.ToUpper(cmds[0])
    handler, ok := handlers[firstToken]
    if ok {
        return handler(h.Engine, cmds)
    } else {
        return errorResponse
    }
}
