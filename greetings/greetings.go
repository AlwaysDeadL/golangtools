// package greetings
package main

import "fmt"

/*
*
main方法只在 package main 中， 其他包不能用 main, 否则程序识别不了主程序
*/
func main() {
	fmt.Println("output", HelloWorld())
}

func HelloWorld() string {
	return "HelloWorld"
}
