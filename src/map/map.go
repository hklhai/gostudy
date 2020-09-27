package main

import "fmt"

func main() {

	var m1 map[string]string
	m1 = make(map[string]string)
	m1["a"] = "123"
	m1["b"] = "12345"

	m2 := make(map[string]string)
	m2["a"] = "123"
	m2["b"] = "12345"

	m3 := map[string]string{
		"a": "123",
		"b": "12345",
	}
	m3["c"] = "12345"

	// 查找键值是否存在
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

}
