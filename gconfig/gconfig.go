package gconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//配置文件 版本1.0
type I struct{}

// Config2Map json配置文件转成Map
func (g *I) Config2Map(file string) (config map[string]interface{}, err error) {
	config = make(map[string]interface{}, 0)

	// 判断文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	// 读取文件内容
	r, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	if err := json.Unmarshal(r, &config); err != nil {
		return nil, err
	}

	return
}

// Config2Struct json配置文件转成结构体
func (g *I) Config2Struct(file string, st interface{}) (err error) {
	// 判断文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return err
	}

	// 读取文件内容
	r, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Unmarshal
	if err := json.Unmarshal(r, &st); err != nil {
		return err
	}

	return
}
