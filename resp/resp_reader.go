package resp

import (
	"bufio"
	"fmt"
	"io"
)

type RESPReader struct {
	*bufio.Reader
}

func NewRESPReader(reader io.Reader) *RESPReader {
	return &RESPReader{
		Reader: bufio.NewReader(reader),
	}

}

func (r *RESPReader) ReadLine() (string, error) {
	line, err := r.Reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("cannot read a line. err: %+v", err)
	}
	if len(line) == 0 {
		return "", fmt.Errorf("line is empty. no way!!!")
	}

	if line[len(line)-2] != '\r' {
		return "", fmt.Errorf("line termination should be '\\r\\n'")
	}
	return line[:len(line)-2], nil
}
