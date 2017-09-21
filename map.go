package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	var Capitalmap map[string]string
	// 创建集合
	Capitalmap = make(map[string]string)
	/* map 插入 key-value 对，各个国家对应的首都 */
	Capitalmap["a"] = "aa"
	Capitalmap["b"] = "bb"
	Capitalmap["c"] = "cc"
	/* 使用 key 输出 map 值 */
	for country := range Capitalmap {
		fmt.Println(country, Capitalmap[country])
	}
	captial, ok := Capitalmap["dd"]
	if ok {
		fmt.Println(captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

}
