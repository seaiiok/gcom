package gcache

import (
	"bytes"
	"encoding/binary"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB
var once sync.Once

type cache struct {
}

func (this *cache) Get(k string) int {
	v, err := db.Get([]byte(k), nil)
	if err != nil {
		return 0
	}
	if len(v) == 0 {
		return 0
	}
	return bytesToInt(v)
}

func (this *cache) Set(k string, v int) {
	db.Put([]byte(k), intToBytes(v), nil)
}

func (this *cache) Delete(k string) {
	db.Delete([]byte(k), nil)
}

func (this *cache) GetMap() (map[string]int, error) {
	m := make(map[string]int)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		m[string(key)] = bytesToInt(value)
	}
	iter.Release()
	err := iter.Error()
	return m, err
}

func (this *cache) Close() {
	defer db.Close()
}

func New(path string) *cache {
	once.Do(func() {
		db, _ = leveldb.OpenFile(path, nil)
	})
	return &cache{}
}

//整形转换成字节
func intToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
