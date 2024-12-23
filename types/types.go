package types

import "encoding/json"

type Animal string
type Person string

type Teacher struct {
	Attribute Person
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
	return Teacher{
		Attribute: Person(name),
		Name:      name,
		age:       0,
		PhoneNum:  138,
	}
}
