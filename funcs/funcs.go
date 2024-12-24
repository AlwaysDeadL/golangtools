package funcs

import (
	"fmt"
	"go/types"
)

type Pointer struct {
	X int32
	Y int32
}

//函数声明包括函数名、形式参数列表、返回值列表（可省略，也可以多返回结果）以及函数体
//func name(parameter-list) (result-list) {
//	body
//}

// 小写的函数，包外部不可访问
func name(name string) string {
	return name
}

// Name 大写的Name 才是对外暴露的函数
func Name(name string) string {
	return name
}

// GetXy 可以不指定返回值
func GetXy(pointer types.Pointer) {
	fmt.Println("executed")
}

// GetXyByReturn 返回一个对象, 参数是形参， 修改后的结构不会修改入参
func GetXyByReturn(pointer Pointer) Pointer {
	fmt.Println("executed")
	pointer.X = 100
	return pointer
}

// GetXyByReturn1 返回一个对象，参数是只传递，
func GetXyByReturn1(pointer *Pointer) *Pointer {
	fmt.Println("executed")
	pointer.X = 100
	return pointer
}

// GetPointerXy 返回多个返回值
func GetPointerXy(pointer Pointer) (int32, int32) {
	fmt.Println("executed")
	return pointer.X, pointer.Y
}

// GetPointerXyByName 可以为多个返回值指定名称
func GetPointerXyByName(pointer Pointer) (X int32, Y int32) {
	fmt.Println("executed")
	return pointer.X, pointer.Y
}

// GetPointerXyByName1 可以为多个返回值指定名称, 前面的类型如果没有指定名称，最后一个必须指定
func GetPointerXyByName1(pointer Pointer) (int32, Y int32) {
	fmt.Println("executed")
	return pointer.X, pointer.Y
}

// GetPointerXyByName1 可以为多个返回值指定名称, 这样写就会报错
//func GetPointerXyByName2(pointer Pointer) (X int32, int32) {
//	fmt.Println("executed")
//	return pointer.X, pointer.Y
//}

// 函数作为参数
func GetPointerByFunc(pointer func(int32, int32) Pointer) Pointer {
	return pointer(10, 10)
}
