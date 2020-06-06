package cache

import (
	"encoding/json"
	"time"

	"github.com/InsideCI/nego/src/models"
	"github.com/dgraph-io/badger/v2"
)

type BadgerRepository struct {
	cache *badger.DB
}

func NewBadgerRepository() *BadgerRepository {

	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))

	if err != nil {
		panic("couldn't initiate cache server" + err.Error())
	}

	return &BadgerRepository{
		cache: db,
	}
}

//SaveByIDKey uses a simple string ID for caching
func (r *BadgerRepository) SaveByIDKey(id string, model interface{}) error {

	byteModel, err := json.Marshal(model)
	if err != nil {
		return err
	}

	err = r.cache.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(id), byteModel).WithTTL(time.Hour * 12)
		return txn.SetEntry(entry)
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *BadgerRepository) SaveByModelKey(key models.QueryParams, model interface{}) error {

	byteKey, err := json.Marshal(key)
	if err != nil {
		return err
	}

	byteModel, err := json.Marshal(model)
	if err != nil {
		return err
	}

	err = r.cache.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(byteKey, byteModel).WithTTL(time.Hour * 12)
		return txn.SetEntry(entry)
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *BadgerRepository) Get(key string) ([]byte, error) {
	var outByte []byte

	err := r.cache.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		outByte, err = item.ValueCopy(nil)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return outByte, nil
}
