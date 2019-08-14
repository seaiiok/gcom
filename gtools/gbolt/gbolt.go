package gbolt

import (
	"sync"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var once sync.Once

type cache struct {
	path   string
	bucket string
}

func New(path, bucket string) *cache {
	this := &cache{
		path:   path,
		bucket: bucket,
	}

	once.Do(func() {
		db = &bolt.DB{}
	})

	if db.Stats().FreeAlloc == 0 {
		db, _ = bolt.Open(path, 0600, nil)
	}

	// Start a writable transaction.
	tx, err := db.Begin(true)
	if err != nil {

	}
	defer tx.Rollback()

	// Use the transaction...
	_, err = tx.CreateBucketIfNotExists([]byte(bucket))
	if err != nil {

	}

	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {

	}
	return this
}

func (this *cache) Get(key string) []byte {
	tx, err := db.Begin(true)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(this.bucket))
	if err != nil {
		return nil
	}

	vb := b.Get([]byte(key))

	if err := tx.Commit(); err != nil {
		return nil
	}

	if len(vb) == 0 {
		return nil
	}

	return vb
}

func (this *cache) Set(key string, value []byte) {
	tx, err := db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(this.bucket))
	if err != nil {
		return
	}

	b.Put([]byte(key), value)

	if err := tx.Commit(); err != nil {
		return
	}
}

func (this *cache) Delete(k string) {
	tx, err := db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(this.bucket))
	if err != nil {
		return
	}

	b.Delete([]byte(k))

	if err := tx.Commit(); err != nil {
		return
	}
}

func (this *cache) GetMap() (list map[string][]byte, err error) {
	list = make(map[string][]byte)
	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(this.bucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			list[string(k)] = v
		}

		return nil
	})
	return
}

func (this *cache) Close() {
	if db != nil {
		db.Close()
	}
}
