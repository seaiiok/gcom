package gcmd

import "testing"

func TestExecCommand(t *testing.T) {
	ExecCommand("chcp", "65001")
	output := ExecCommand("cmd", "/C","VER")
	t.Log(output)
}

func TestPrint(t *testing.T) {
	for i := 0; i < 15; i++ {
		Println(i, "终端颜色-",i)
	}
}
