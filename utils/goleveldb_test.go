package utils

import (
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

var (
	key    = "test"
	value  = "test"
	dbName = "test_db"
)

func TestWriteDB(t *testing.T) {
	db, err := initDB(dbName)
	require.NoError(t, err)
	batch := new(leveldb.Batch)
	batch.Put([]byte(key), []byte(value))
	err = db.Write(batch, nil)
	require.NoError(t, err)
}

func TestReadListDB(t *testing.T) {
	db, err := initDB(dbName)
	require.NoError(t, err)

	var privateKey []string
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		value := iter.Value()
		privateKey = append(privateKey, string(value))
	}
	iter.Release()
	t.Log(privateKey)
	err = iter.Error()
	require.NoError(t, err)
}

func TestDeleteByKey(t *testing.T) {
	db, err := initDB(dbName)
	require.NoError(t, err)
	err = db.Delete([]byte(key), nil)
	require.NoError(t, err)
}
