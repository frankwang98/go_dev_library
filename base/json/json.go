package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	/* 解析json */
	// 要解析的 JSON 数据
	jsonData := []byte(`{"name":"Alice","age":25}`)

	// 创建一个 Person 类型的变量
	var p Person

	// 解析 JSON 数据到 Person 变量
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return
	}

	// 输出解析结果
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)


	/* 写入json */
	// // 创建一个 Person 对象
	// p := Person{Name: "Bob", Age: 30}

	// // 将 Person 对象转换为 JSON 格式
	// jsonData, err := json.Marshal(p)
	// if err != nil {
	// 	fmt.Println("生成 JSON 出错:", err)
	// 	return
	// }

	// // 输出生成的 JSON 数据
	// fmt.Println(string(jsonData)) 
}
