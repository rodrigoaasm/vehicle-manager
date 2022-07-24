package myleveldb_test

import (
	"demo/external/datasource/myleveldb"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db, errDatabase = myleveldb.NewDatabase("test/" + strconv.Itoa(rand.Intn(100)) + "/")

func TestNewDatabase(t *testing.T) {
	require.Nil(t, errDatabase, "should create a local database")
	assert.NotNil(t, db.Data, "should create a database")
	assert.NotNil(t, db.Index, "should create an index database")
}

func TestClose(t *testing.T) {
	db.Close()
}
