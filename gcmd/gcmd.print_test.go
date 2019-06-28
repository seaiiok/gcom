package gcmd

import "testing"

func TestPrintColor(t *testing.T) {
	ii := I{}
	for i := 0; i < 15; i++ {
		ii.Println(i, "嗨喽,Seaii！", "Color:", i)
	}
}
