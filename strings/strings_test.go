package strings

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	//字符串是不可变的
	//不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，
	//1. 这使得复制任何长度的字符串代价是低廉的。
	//2. 同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的。
	//在这两种情况下都没有必要分配新的内存。
	s := "left foot"
	r := s
	s += ", right foot"
	fmt.Println(r) // left foot
	fmt.Println(s) // left foot, right foot

	if !HasPrefix("hello world", "hello") {
		t.Errorf("not match prefix.")
	}

	if !HasSuffix("hello world.sh", "sh") {
		t.Errorf("not match suffix.")
	}

	if !Contains("hello world.sh", "or") {
		t.Errorf("not match suffix.")
	}
}
