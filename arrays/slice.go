package arrays

import (
	"fmt"
)

func SliceSamples() {
	//默认初始化下标为: 0
	weeksOri := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(weeksOri[6], len(weeksOri), cap(weeksOri)) //Sunday, 长度 7， 容量 7

	//可以初始化的时候指定第一个元素的下标，那么后续元素不指定下标的话，会自动递增如下, 而weeks[0]的元素则会根据类型初始化为 ""， 如果是int, 那么就是 0
	weeks := [...]string{1: "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(weeks[0], weeks[0] == "")         // ""空串, true
	fmt.Println(weeks[7], len(weeks), cap(weeks)) //Sunday， 长度 8， 容量 8
	fmt.Println(weeks)

	//但是如果指定的是最后一个元素的下标，那么前面的元素不会递减， 还是会按默认的 0，1，2，递增，然后, [7],[8],[9] 是 ""
	weeks2 := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", 10: "Sunday"}
	fmt.Println(weeks2[0], weeks2[7], weeks2[8], weeks2[9], weeks2[10], len(weeks2), cap(weeks2))

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)

	years := make([]int, 1, 3) // 切片类型，1是长度， 3是容量
	years[0] = 2012
	//years[1] = 2013
	//years[2] = 2014
	//years[3] = 2015 //这样会报错，因为slice指定了长度为 1， 直接使用下标添加会报错，要扩展slice,需要用append
	fmt.Println(years, len(years), cap(years))
	years = append(years, 2013, 2014, 2015)
	fmt.Println(years, len(years), cap(years))

	summers := months[3:6]
	fmt.Println(summers)
	springs := months[:3]
	fmt.Println(springs)
	winters := months[9:]
	fmt.Println(winters)

	/*
		数组与切片的区别:
			1. 数组
				1.1 具有固定长度的序列, 长度属于类型的一部分，定义时需要指定长度，一旦声明，数组的大小不可更改
				1.2 数组的长度跟容量相同
				1.3 存储的是实际的数据，分配在栈上
				1.4 赋值会拷贝整个数组
				1.5 开销大，不建议使用
			2. 切片
				2.1 可以动态改变长度大小, 它是一个对底层数组的引用，允许对该数组某部分进行更改
				2.2 长度是切片实际的元素数量，容量是从切片的起始位置到底层数组的末尾的元素数量
				2.3 可以使用append来对slice动态扩容
				2.4 只是对底层数组的引用，存储指针，长度和容量, 分配在堆上
				2.5 多个切片可以共享底层数据
				2.6 开销小，因为只传递引用，一般都用slice
		数组与切片的内存分配
			1. 数组
				1.1 静态大小： 在编译时确定，编译器可以直接分配固定大小的内存
				1.2 存储位置:
						1.2.1 如果是局部变量， 数组的内容通常分配在栈上.
						1.2.2 如果是全局变量， 数组内容分配在全局静态存储区
						1.2.3 如果动态初始化(new, make), 数组内容在堆上
				1.3 开销:
						1.3.1 数组在栈上: 分配和访问速度快，内存大小有限，超过
			2. 切片
				2.1 动态大小: 切片本身只是一个结构体，包含数组的指针，长度，容量
				2.2 存储位置: 切片本身存储在栈上，但是引用的数组数据可能在 栈上 (小容量数组可能分配在栈上)，也可能在堆上

	*/

}
