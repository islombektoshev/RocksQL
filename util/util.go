package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const TestDirName = "rocksql-temp-test"

func PrepareTestDirPath() string {
	return os.TempDir() + "/" + TestDirName + "/" + strconv.FormatInt(time.Now().UnixMicro(), 10)
}

func MustMakeDir(path string) {
	err := os.MkdirAll(path, os.ModeDir|os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(fmt.Errorf("cannot make dir. path: %s, err: %+v", path, err))
	}
}

func MustPrepareTestDir() string {
	var path = PrepareTestDirPath()
	MustMakeDir(path)
	return path
}

func MustDeleteTestDir(path string) {
	if strings.Contains(path, TestDirName) {
		err := os.RemoveAll(path)
		if err != nil {
			panic(fmt.Errorf("unable to delete dir. path: %s. err: $%+v", path, err))
		}
	} else {
		panic("provide path  to delete is not belong to test")
	}
}

func Tokenize(line string) ([]string, error) {
	var start = 0
	var inDQ = false
	var tokens []string = []string{}
	for i, r := range line {
		if r == '"' {
			if inDQ {
				tokens = append(tokens, line[start+1:i])
				start = i + 1
				inDQ = false
			} else {
				if start != i {
					tokens = append(tokens, line[start:i])
				}
				inDQ = true
				start = i
			}
		}
		if r == ' ' && !inDQ {
			//log.Printf("start %d, i: %d, r: %v\n", start, i, r)
			if start != i {
				token := line[start:i]
				//log.Printf("token: %s", token)
				tokens = append(tokens, token)
			}
			start = i + 1
		}
	}
	if start < len(line) {
		tokens = append(tokens, line[start:])
	}
	if inDQ {
		return nil, fmt.Errorf("not closing quote")
	}
	return tokens, nil
}

func RandString(rn *rand.Rand, maxBound int, notAllowed []byte) string {
	count := rn.Intn(maxBound)
	arr := make([]byte, count)
	for i := 0; i < count; {
		bayt := byte(rn.Intn(1 << 8))
		if bytes.Contains(notAllowed, []byte{bayt}) {
			continue
		}
		arr[i] = bayt
		i++
	}
	return string(arr)
}
