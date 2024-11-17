package resp_test

import (
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/islombektoshev/RocksQL/resp"
	"github.com/stretchr/testify/assert"
)

func TestRESPReader_ReadLine_Trivial_Ok(t *testing.T) {

	pipeReader, pipeWriter := io.Pipe()
	reader := resp.NewRESPReader(pipeReader)
	go func() {
		line, err := reader.ReadLine()
		assert.NoError(t, err)
		assert.Equal(t, line, "aline")
	}()
	n, err := pipeWriter.Write([]byte("aline\r\n"))
	assert.NoError(t, err)
	assert.Equal(t, n, 7)
}

func TestRESPReader_ReadLine_Emptry_OK(t *testing.T) {

	pipeReader, pipeWriter := io.Pipe()
	reader := resp.NewRESPReader(pipeReader)
	go func() {
		line, err := reader.ReadLine()
		assert.NoError(t, err)
		assert.Equal(t, line, "")
	}()
	n, err := pipeWriter.Write([]byte("\r\n"))
	assert.NoError(t, err)
	assert.Equal(t, n, 2)
}

func TestRESPReader_ReadLine_Error_Only_LF(t *testing.T) {

	pipeReader, pipeWriter := io.Pipe()
	reader := resp.NewRESPReader(pipeReader)
	go func() {
		line, err := reader.ReadLine()
		assert.Equal(t, line, "")
		assert.NotNil(t, err)
	}()
	n, err := pipeWriter.Write([]byte("aline\n"))
	assert.NoError(t, err)
	assert.Equal(t, n, 6)
}

func TestRESPReader_ReadLine_Error_Ok_Seperate_writing(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	reader := resp.NewRESPReader(pipeReader)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		line, err := reader.ReadLine()
		fmt.Printf("line: %s, err: %+v", line, err)
		assert.NoError(t, err)
		assert.Equal(t, line, "aline")
	}()

	n, err := pipeWriter.Write([]byte("aline"))
	assert.NoError(t, err)
	assert.Equal(t, n, 5)
	time.Sleep(10 * time.Millisecond)

	n, err = pipeWriter.Write([]byte("\r"))
	assert.NoError(t, err)
	assert.Equal(t, n, 1)

	n, err = pipeWriter.Write([]byte("\n"))
	assert.NoError(t, err)
	assert.Equal(t, n, 1)
	wg.Wait()
}

func TestPipedReader(t *testing.T) {
	reader, writer := io.Pipe()
	respReader := resp.NewRESPReader(reader)

	go func() {
		fmt.Println("2 read line test")
		line, err := respReader.ReadLine()
		assert.NoError(t, err)
		assert.Equal(t, "test", line)

		fmt.Println("4 read nextline")
		line, err = respReader.ReadLine()
		assert.NoError(t, err)
		assert.Equal(t, "nextline", line)

	}()

	fmt.Println("1 write test")
	_, err := writer.Write([]byte("test\r\n"))
	assert.NoError(t, err)

	fmt.Println("3 write nextline")
	_, err = writer.Write([]byte("nextline\r\n"))
	assert.NoError(t, err)

	fmt.Println("complete")
}
