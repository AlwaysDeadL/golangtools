package types

import "encoding/json"

type Shape string

// Animal Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针
type Animal struct {
	Shape //golang特有的匿名成员,
}
type Person string

type Teacher struct {
	Attribute Person
	Animal    //golang特有的匿名成员,
	Name      string
	age       int // 小写的字段 = private, json.Marshell的时候不会序列化，但是这个字段可读, 如果想保持private, 但是要序列化出来需要自定义序列化方法
	//PhoneNum  int `json: "phone_num"` //注意， 不能有空格，否则自定义序列化字段不生效
	PhoneNum int `json:"phone_num"`
}

// MarshalJSON Custom MarshalJSON method for Teacher struct
func (t Teacher) MarshalJSON() ([]byte, error) {
	// Create a new struct to hold the serialized data
	type Alias Teacher // Create an alias to avoid recursion
	return json.Marshal(&struct {
		*Alias
		Age int `json:"age"` // Include `age` in the JSON with the custom key "age"
	}{
		Alias: (*Alias)(&t),
		Age:   t.age, // Explicitly include the unexported `age` field
	})
}

func GetTeacher(name string) Teacher {
	teacher := Teacher{
		Attribute: Person(name),
		Name:      name,
		age:       0,
		PhoneNum:  138,
	}

	//其中匿名成员Animal和Shape有自己的名字——就是命名的类型名字
	//得益于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径：
	teacher.Animal.Shape = "monkey"
	teacher.Shape = "tangle"
	return teacher
}
