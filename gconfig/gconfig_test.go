package gconfig

import "testing"

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
	g := &I{}
	testfile := "./config_test.json"

	confmap, err := g.Config2Map(testfile)
	if err != nil {
		t.Log(err)
	}
	t.Log(confmap)

	c := &conf{}
	err = g.Config2Struct(testfile, &c)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}