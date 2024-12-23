package maps

import (
	"encoding/json"
	"fmt"
)

// MapSamples just used for test map usage
func MapSamples() {
	//声明一个map, 使用内置的make函数
	var localMap = make(map[string]int)
	localMap["jack"] = 1
	localMap["Maple"] = 2
	data, _ := json.Marshal(localMap)
	fmt.Println(string(data))

	//直接使用map实例化
	var localMap2 = map[string]string{}
	localMap2["hello"] = "world"
	localMap2["good"] = "bye"
	data, _ = json.Marshal(localMap2)
	fmt.Println(string(data))

	//删除数据, 使用内置函数
	delete(localMap, "jack")

	//遍历数据
	for key, value := range localMap2 {
		fmt.Println("key:", key, "value:", value)
	}

	//声明一个空map, 可以查，删除，取长度，遍历，但是不能赋值，要赋值必须得先初始化
	var nilMap map[string]int
	nilMapLength := len(nilMap)
	fmt.Println("nilMapLength:", nilMapLength)
	delete(nilMap, "what")

	//会报错 panic: assignment to entry in nil map [recovered]
	//	panic: assignment to entry in nil map
	//nilMap["add"] = 1

	nilMap = make(map[string]int)
	nilMap["add"] = 1
	fmt.Println(nilMap)

}
