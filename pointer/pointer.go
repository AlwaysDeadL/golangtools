package pointer

import "fmt"

// HelloPointer /**
//  1. 一个变量对应一个保存了变量对应类型值的内存空间
//  1. 一个指针的值是另一个变量的地址，即指针存储的是变量在内存的地址
//  2. 不是每个值都会对应一个内存地址， 但是对于每个变量必定有对应的内存地址。通过指针可以读或更新变量的值，而不需要知道变量的名称
//  3. 如果用 var x int = 1 声明了一个x变量，且变量值为 1， 那么 &x 会产生一个指向x变量的指针，指针对应的数据类型是 *int.
//     指针被称为 "指向int类型的指针"
//     如果该指针名字为p, 那么可以说 "p（&x）指针指向变量x", 或者说 "p (&x)指针保存了变量x的内存地址"
//     同时， "*p"表达式则对应 "p指针指向的x变量的值"
func HelloPointer() {
	x := 1
	// &x 为指针，且值为 x变量的内存地址
	p := &x
	// p = &x, 例如 p = &x =  0x1400000e378
	fmt.Println("p:", p)
	fmt.Println("&x:", &x)
	//&p 会产生一个新的指针，指向p变量(也就是&x)的内存地址， 所以 &p != &x, 他们是两个不同的内存地址
	fmt.Println("&p:", &p) // &p = 0x14000056050
	//*p 指向 x变量的值, 所以*p = x = 1,
	fmt.Println("*p=x=*&x:", *p, x, *&x)
	//*&p = p = &x = 0x1400000e378
	fmt.Println("*&p:", *&p)
	// 这里 *p = 2, 实际上是 *&x = 2, 那么就相当于修改了 x的值， 所以x = 2, p本身还是 = &x，所以还是等于x的值的内存地址
	*p = 2
	fmt.Println("*p = 2:", *p)
	fmt.Println("p:", p)
	fmt.Println("x:", x)
	x = 3
	fmt.Println("x = 3, *p:", *p)
	fmt.Println("**&p = **&&x = x , 所以**&p=3", **&p)
}
