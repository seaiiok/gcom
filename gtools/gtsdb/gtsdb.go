package gtsdb

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var once sync.Once
var db *leveldb.DB

type cache struct {
	items map[string]*List
}

type List struct {
	Key   int64
	Value interface{}
	Next  *List
}

func New(root string) *cache {

	once.Do(func() {
		db, _ = leveldb.OpenFile(root, nil)
	})

	return &cache{
		items: make(map[string]*List, 0),
	}
}

func (this *cache) Set(key string, list *List) {
	b, _ := json.Marshal(list)
	db.Put([]byte(key), b, nil)
}

func (this *cache) Get(key string) (list *List) {
	v, _ := db.Get([]byte(key), nil)
	json.Unmarshal(v, &list)
	return
}

func (this *cache) GetList(prefix string) (items map[string]*List, err error) {
	items = make(map[string]*List)

	iter := db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		v := &List{}
		json.Unmarshal(value, &v)
		items[string(key)] = v
		fmt.Println(v)
	}
	iter.Release()
	err = iter.Error()
	return
}

func (this *cache) Print(m *List) {
	e := m
	for {
		if e == nil {
			break
		}

		fmt.Println("print:", e.Key, e.Value)

		e = e.Next
	}
}

func (this *List) SetList(key int64, value interface{}) (list *List) {
	l := &List{key, value, nil}
	if this != nil {
		l.Next = this
	}
	list = l
	return
}

func (this *List) GetListLimit(s, e int64) (list *List) {
	list = &List{}
	for {
		if this == nil {
			return
		}

		if this.Key >= s && this.Key <= e {
			l := &List{this.Key, this.Value, nil}
			l.Next = list
			list = l
		}

		this = this.Next
	}
}
