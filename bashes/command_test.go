package bashes

import (
	"testing"
)

func TestExecBashFile(t *testing.T) {
	var filePath = "./ls.sh"
	ExecBashFile(filePath)

}
