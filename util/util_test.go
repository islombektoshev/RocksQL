package util_test

import (
	"testing"

	"github.com/islombektoshev/RocksQL/util"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	res, err := util.Tokenize("GET 1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"GET", "1"}, res)

	res, err = util.Tokenize("SET 1 value1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "1", "value1"}, res)

	res, err = util.Tokenize("SET abc value1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "abc", "value1"}, res)

	res, err = util.Tokenize("SET 1 \"salom dunyo\" ")
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "1", "salom dunyo"}, res)

	res, err = util.Tokenize(`SET  1"salom dunyo" `)
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "1", "salom dunyo"}, res)

	res, err = util.Tokenize(`SET"key 1""salom dunyo" `)
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "key 1", "salom dunyo"}, res)

	res, err = util.Tokenize(`"SET""key 1""salom dunyo" `)
	assert.NoError(t, err)
	assert.Equal(t, []string{"SET", "key 1", "salom dunyo"}, res)

	res, err = util.Tokenize(``)
	assert.NoError(t, err)
	assert.Equal(t, []string{}, res)

	res, err = util.Tokenize(`    `)
	assert.NoError(t, err)
	assert.Equal(t, []string{}, res)

	res, err = util.Tokenize(`GET "1`)
	assert.NotNil(t, err)

}
