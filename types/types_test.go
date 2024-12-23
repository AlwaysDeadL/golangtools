package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTypes(t *testing.T) {
	var teacher = GetTeacher("Jack")
	fmt.Println(teacher)

	fmt.Println(teacher.age)

	jsonData, _ := json.Marshal(teacher)
	fmt.Println(string(jsonData))

}
