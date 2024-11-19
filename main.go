package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/server"
	"github.com/islombektoshev/RocksQL/util"
	"github.com/linxGnu/grocksdb"
)

const DefaultDbPath = "/tmp/rocksql/db"

var (
	dbPath string
	port   int
	host   string
)

func main() {
	flag.StringVar(&dbPath, "dbpath", "rocksql-data", "Please set you some dbpath to run")
	flag.StringVar(&host, "host", "localhost", "Please set sever host to run")
	flag.IntVar(&port, "port", 6666, "Pease set server port to run")

	flag.Parse()

	var opt = grocksdb.NewDefaultOptions()
	opt.SetCreateIfMissing(true)
	defer opt.Destroy()

	var db, err = grocksdb.OpenDb(opt, dbPath)
	if err != nil {
		panic(fmt.Errorf("cannot open db: %+v", err))
	}
	defer db.Close()

	eng := engine.NewEngine(db)
	svr := server.NewServer(eng, fmt.Sprintf("%s:%d", host, port))
	err = svr.StartListening()
	if err != nil {
		fmt.Printf("cannot listen port 6666, err: %+v", err)
	}
	svr.AcceptContinously(context.Background())

	// Loop(eng)
}

func Loop(eng engine.Engine) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			log.Fatal("ERR: {}", err)
		}
		tokens, err := util.Tokenize(line)
		if err != nil {
			log.Printf("ERR: %v", err)
			continue
		}
		if len(tokens) < 1 {
			continue
		}

		var firstToken = strings.ToUpper(tokens[0])
		switch firstToken {
		case "GET":
			if len(tokens) != 2 {
				fmt.Println("Wrong format GET. Example: `GET 1`")
				continue
			}
			var key = tokens[1]
			// fmt.Printf("\nDEBUG GET '%s'", key);
			var res = eng.Get([]byte(key))
			if res.Err != nil {
				log.Printf("Error occurding during GET: %+v", err)
			}
			fmt.Println(string(res.Data))
			continue
		case "PUT":
			if len(tokens) != 3 {
				fmt.Println("Wrong format PUT. Example: `PUT 1 value`")
				continue
			}
			var key = tokens[1]
			var val = tokens[2]
			//fmt.Printf("\nDEBUG PUT '%s' '%s'", key, val);
			var res = eng.Put([]byte(key), []byte(val))
			if res.Err != nil {
				log.Printf("Error occurding during GET: %+v", err)
			} else {
				fmt.Println("OK")
			}
			continue

		case "EXIT":
			return
		default:
			fmt.Println("unknown command, ex: GET, PUT")
			continue
		}
	}
}
