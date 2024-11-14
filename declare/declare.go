package declare

import "fmt"

/*
DoDeclareSamples docs
有三种方式声明变量
1. var: 可以声明变量并赋值， 如： var a = 1， 或者先声明再赋值，如: var a int; a = 2
2. := : 这个表示声明并赋值， 且会自动类型推断，不需要声明变量的类型。 如 a := 1
3. const: 常量声明

三种写法都可以同时声明多个变量， 如:
var a,b,c = 1,2,3
a,b,c := 1,2,3
const a,b,c = 1,2,3
 */
const name = "The"
func DoDeclareSamples() {
	var a string
	a = "m"
	//可以改变值
	a = "c"
	//a = 1这里会报错，因为"a"声明了类型string
	//a = 1

	var b int = 1
	var c = 1
	var d int; d = 2
	var e,f,g = 1,2,3
	fmt.Println(a,b,c,d,e,f,g)

	const h,i = 1,2
	fmt.Println(h,i)
	//这会报错，常量只能声明一次
	//h = "n"

	j := 1
	//j = "A" 这里也会报错，因为:= 等价于 var j int = 1, 此时j就是int类型，再赋值字符串就会报错
	//j = "A"
	k,l := 1,2
	fmt.Println(j, k, l)
}


func DoAdd(a int, b int) int {
	return a + b
}