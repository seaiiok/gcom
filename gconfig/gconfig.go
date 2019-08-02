package gconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config2Map json配置文件转成Map
func Config2Map(file string) (config map[string]interface{}, err error) {
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

// Config2MapStr json配置文件转成Map String
func Config2MapStr(file string) (config map[string]string, err error) {
	config = make(map[string]string, 0)

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
func Config2Struct(file string, st interface{}) (err error) {
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

// Config2ListMap json配置文件转成Map
func Config2ListMap(file string) (map[string]interface{}, error) {
	config := make(map[string]interface{}, 0)

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

	m := make(map[string]interface{})

	for k, v := range config {
		// tv := reflect.TypeOf(v)
		// vv := reflect.ValueOf(v)
		switch v.(type) {
		case map[string]interface{}:
		
		case string:
			m[k] = v.(string)
		case int:
			m[k] = v.(int)
		case float64:
			m[k] = fmt.Sprintf("%s", v)
			fmt.Println("age")
		}
	}

	return m, nil
}
