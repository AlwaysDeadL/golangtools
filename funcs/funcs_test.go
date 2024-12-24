package funcs

import (
	"fmt"
	"testing"
)

func TestGetXy(t *testing.T) {
	//result := GetXy(types.Pointer{})
	//fmt.Println(result)
	var pointer = Pointer{X: 10, Y: 10}

	x, y := GetPointerXy(pointer)
	fmt.Println(x, y)
	x, y = GetPointerXyByName(pointer)
	fmt.Println(x, y)
	x, y = GetPointerXyByName1(pointer)
	fmt.Println(x, y)

	result := GetXyByReturn(pointer)
	fmt.Println(result, pointer)

	newResult := GetXyByReturn1(&pointer)
	fmt.Println(newResult, pointer)

	pointer = GetPointerByFunc(func(i int32, i2 int32) Pointer {
		return Pointer{X: 10, Y: 10}
	})
	fmt.Println("pointer: ", pointer)
}

func TestPingError(t *testing.T) {
	_, _ = ReturnNil1("jack.com")
	_ = WaitForServer("jack.com")
	PingError()
	ignoreError()
	FatalAndExit(11)
}
