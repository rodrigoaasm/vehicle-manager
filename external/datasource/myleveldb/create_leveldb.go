package myleveldb

import "github.com/syndtr/goleveldb/leveldb"

type Database struct {
	Data  *leveldb.DB
	Index *leveldb.DB
}

func NewDatabase(path string) (*Database, error) {
	data, err := leveldb.OpenFile(path+"data", nil)
	if err != nil {
		return nil, err
	}

	index, err := leveldb.OpenFile(path+"index", nil)
	if err != nil {
		data.Close()
		return nil, err
	}

	db := Database{
		Data:  data,
		Index: index,
	}

	return &db, nil
}

func (db Database) Close() {
	db.Data.Close()
	db.Index.Close()
}
