package enums_test

import (
	"fmt"
	"github.com/AlwaysDeadL/golangtools/enums"
	"testing"
)

func TestEnums(t *testing.T) {

	var monday = enums.Wednesday
	fmt.Println(monday)

	var autumn = enums.Autumn
	fmt.Println(autumn)
}
