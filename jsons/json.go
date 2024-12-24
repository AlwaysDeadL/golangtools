package jsons

import (
	"encoding/json"
	"fmt"
	"github.com/AlwaysDeadL/golangtools/types"
)

type Course struct {
	Type    string `json:"courseName"`
	Price   float32
	Content string
	Teacher types.Teacher `json:"coach"`
}

func JsonExample() {
	course := Course{
		Type:    "java",
		Price:   0.1,
		Content: "java content",
		Teacher: types.Teacher{
			Attribute: types.Person("jack"),
			Name:      "jack",
			//age:       0,
			PhoneNum: 138,
		},
	}

	data, err := json.Marshal(course)
	if err == nil {
		fmt.Println(string(data))
	}
}
