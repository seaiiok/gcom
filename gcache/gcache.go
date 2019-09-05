package gcache

import (
	"sync"
	"time"
)

const (
	TagQualityGood = 192
)

type Item struct {
	ChannelName             string
	ChannelDriver           string
	DeviceName              string
	TagAddress              string
	TagDataType             string
	TagReadWriteAccess      string
	TagScanRateMilliseconds string
	TagDescription          string
	TagTimeStamp            int64
	TagQuality              int64
}

type cache struct {
	items          map[string]Item
	itemExpiration int64
	mu             sync.RWMutex
}

func (this *cache) Set(key string, value Item) {

	//判断缓存数据质量
	if value.TagQuality != TagQualityGood {
		return
	}

	//判断缓存是否过期
	if time.Now().UnixNano()-value.TagTimeStamp > this.itemExpiration {
		return
	}

	this.items[key] = value
}

func (this *cache) Get(key string) (Item, bool) {
	this.mu.RLock()

	//判断缓存中是否存在
	item, found := this.items[key]
	if !found {
		this.mu.RUnlock()
		return Item{}, false
	}

	//判断缓存质量
	if item.TagQuality != TagQualityGood {
		this.mu.RUnlock()
		return Item{}, false
	}

	//判断缓存是否过期
	if time.Now().UnixNano()-item.TagTimeStamp > this.itemExpiration {
		this.mu.RUnlock()
		return Item{}, false
	}

	this.mu.RUnlock()
	return item, true
}
