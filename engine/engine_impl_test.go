package engine_test

import (
	"testing"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/util"
	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	var testDir = util.MustPrepareTestDir()
	defer util.MustDeleteTestDir(testDir)

	var db, err = engine.OpenDB(testDir)
	assert.NoError(t, err, "OpenDb")

	var engine = engine.NewEngine(db)
	res := engine.Put([]byte("1"), []byte("val2"))
	assert.NoError(t, res.Err)

	dataRes := engine.Get([]byte("1"))
	assert.NoError(t, dataRes.Err)
}
