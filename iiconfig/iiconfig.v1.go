package iiconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//V1 版本v1.0
type V1 struct {
}

//New ...return V1
func New() *V1 {
	return &V1{}
}

func (c *V1) primeCacheFromFile(file string) (*map[string]interface{}, error) {
	// 判断文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	// Read file
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var config map[string]interface{}
	if err := json.Unmarshal(raw, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
