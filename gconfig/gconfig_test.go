package gconfig

import (
	"testing"
)

func TestConfig(t *testing.T) {

	testfile := "./config_test.json"

	err := UpdateConfig(testfile)
	t.Log(err)

	r := GetConfig("gender")
	t.Log(r)
}
