package arrays

import (
	"fmt"
	"github.com/AlwaysDeadL/golangtools/enums"
)

type Hello struct {
	a       string
	weekday enums.Weekday
}

func ArraySample() {
	//1.数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成
	//2.因为数组的长度是固定的，因此在Go语言中很少直接使用数组
	//3.数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。内置的len函数将返回数组中元素的个数
	//4.默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0
	var a [3]int
	a[0] = 4
	a[2] = 5
	fmt.Println(a[1]) //a[i]会初始化为0

	var b [3]string
	fmt.Println(b[1]) //b[1]会初始化为"", 空串
	fmt.Println(b[1] == "")

	var c [3]Hello
	fmt.Println(c) // { 0}, 对象类型的数组也会初始化字段， 那么str="", weekday=0

	var d [3]int = [3]int{1, 2, 3}
	f := [3]int{1, 2}
	fmt.Println(d)
	fmt.Println(f)

	//可以用 [...] 不声明数组的具体长度，而是根据实例化的值来决定数组的长度
	var g = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(g)
}
