package myleveldb

import "github.com/syndtr/goleveldb/leveldb"

func CreateLevelDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)

	return db, err
}
