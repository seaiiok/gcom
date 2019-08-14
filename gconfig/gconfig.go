package gconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

const prekey = "MES-Collector-20181001 09:00:00"

// parseConfig ...
func parseConfig(file string) ([]byte, error) {

	// 判断文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	// 读取文件内容
	r, err := ioutil.ReadFile(file)

	if err != nil {

		return nil, err
	}
	return r, nil
}

func UpdateConfig(file string) error {
	b, err := parseConfig(file)
	if err != nil {
		return err
	}

	config := make(map[string]interface{}, 0)
	if err := json.Unmarshal(b, &config); err != nil {
		return err
	}

	for k, v := range config {
		value := ""
		switch v.(type) {
		case string:
			value = v.(string)
		case bool:
			value = strconv.FormatBool(v.(bool))
		case float64:
			value = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		default:
			value = ""
		}

		if err := os.Setenv(prekey+k, value); err != nil {
			return err
		}
	}
	return nil
}

func Set(k, v string) {
	os.Setenv(prekey+k, v)
}

func Get(k string) string {
	return os.Getenv(prekey + k)
}
