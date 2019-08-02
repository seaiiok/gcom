package gconfig

import (
	"testing"
)

type conf struct {
	Name  string
	Age   int
	Score score
}

type score struct {
	Chinese float64
	English float64
}

func TestConfig(t *testing.T) {

	testfile := "./config_test.json"

	confmap, err := Config2ListMap(testfile)
	if err != nil {
		t.Log(err)
	}

	t.Log(confmap)

	// c := &conf{}
	// err = Config2Struct(testfile, &c)
	// if err != nil {
	// 	t.Log(err)
	// }
	// t.Log(c)
}
